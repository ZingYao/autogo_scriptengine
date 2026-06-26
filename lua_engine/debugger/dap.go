package debugger

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"
	"sync"
)

// DAPSession 处理 Debug Adapter Protocol 的基础请求。
type DAPSession struct {
	debugger *Debugger
	in       *bufio.Reader
	out      io.Writer
	mu       sync.Mutex
	seq      int
}

type dapMessage struct {
	Seq        int             `json:"seq,omitempty"`
	Type       string          `json:"type"`
	Command    string          `json:"command,omitempty"`
	Event      string          `json:"event,omitempty"`
	RequestSeq int             `json:"request_seq,omitempty"`
	Success    bool            `json:"success,omitempty"`
	Message    string          `json:"message,omitempty"`
	Body       interface{}     `json:"body,omitempty"`
	Arguments  json.RawMessage `json:"arguments,omitempty"`
}

// NewDAPSession 创建一个 DAP 会话。
func NewDAPSession(debugger *Debugger, input io.Reader, output io.Writer) *DAPSession {
	return &DAPSession{
		debugger: debugger,
		in:       bufio.NewReader(input),
		out:      output,
	}
}

// Serve 阻塞处理 DAP 请求，直到输入流关闭或出现协议错误。
func (s *DAPSession) Serve() error {
	go s.forwardEvents()
	for {
		msg, err := s.readMessage()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		if msg.Type != "request" {
			continue
		}
		s.handleRequest(msg)
	}
}

func (s *DAPSession) forwardEvents() {
	for event := range s.debugger.Events() {
		switch event.Type {
		case "stopped":
			s.sendEvent("stopped", map[string]interface{}{
				"reason":            string(event.Reason),
				"threadId":          event.ThreadID,
				"allThreadsStopped": true,
				"description":       event.Message,
			})
		case "terminated":
			s.sendEvent("terminated", map[string]interface{}{})
		}
	}
}

func (s *DAPSession) handleRequest(req dapMessage) {
	switch req.Command {
	case "initialize":
		s.sendResponse(req, true, map[string]interface{}{
			"supportsConfigurationDoneRequest": true,
			"supportsSetVariable":              false,
			"supportsEvaluateForHovers":        false,
			"supportsStepBack":                 false,
		}, "")
		s.sendEvent("initialized", map[string]interface{}{})
	case "launch", "attach", "configurationDone":
		s.sendResponse(req, true, map[string]interface{}{}, "")
	case "setBreakpoints":
		s.handleSetBreakpoints(req)
	case "threads":
		s.sendResponse(req, true, map[string]interface{}{
			"threads": []map[string]interface{}{{"id": defaultThreadID, "name": "Lua main"}},
		}, "")
	case "continue":
		s.debugger.Continue()
		s.sendResponse(req, true, map[string]interface{}{"allThreadsContinued": true}, "")
	case "next", "stepIn":
		s.debugger.StepInto()
		s.sendResponse(req, true, map[string]interface{}{}, "")
	case "pause":
		s.debugger.Pause()
		s.sendResponse(req, true, map[string]interface{}{}, "")
	case "disconnect", "terminate":
		s.debugger.Stop()
		s.sendResponse(req, true, map[string]interface{}{}, "")
	case "stackTrace":
		s.handleStackTrace(req)
	case "scopes":
		s.handleScopes(req)
	case "variables":
		s.handleVariables(req)
	default:
		s.sendResponse(req, false, nil, "unsupported command: "+req.Command)
	}
}

func (s *DAPSession) handleSetBreakpoints(req dapMessage) {
	var args struct {
		Source struct {
			Path string `json:"path"`
			Name string `json:"name"`
		} `json:"source"`
		Breakpoints []struct {
			Line int `json:"line"`
		} `json:"breakpoints"`
		Lines []int `json:"lines"`
	}
	if err := json.Unmarshal(req.Arguments, &args); err != nil {
		s.sendResponse(req, false, nil, err.Error())
		return
	}
	file := args.Source.Path
	if file == "" {
		file = args.Source.Name
	}
	lines := args.Lines
	for _, bp := range args.Breakpoints {
		lines = append(lines, bp.Line)
	}
	points := s.debugger.SetBreakpoints(file, lines)
	response := make([]map[string]interface{}, 0, len(points))
	for _, point := range points {
		response = append(response, map[string]interface{}{
			"id":       point.ID,
			"verified": point.Verified,
			"line":     point.Line,
			"source": map[string]interface{}{
				"path": point.File,
			},
		})
	}
	s.sendResponse(req, true, map[string]interface{}{"breakpoints": response}, "")
}

func (s *DAPSession) handleStackTrace(req dapMessage) {
	frame := s.debugger.LastFrame()
	if frame == nil {
		s.sendResponse(req, true, map[string]interface{}{"stackFrames": []interface{}{}, "totalFrames": 0}, "")
		return
	}
	s.sendResponse(req, true, map[string]interface{}{
		"stackFrames": []map[string]interface{}{
			{
				"id":     frame.ID,
				"name":   frame.Name,
				"line":   frame.Position.Line,
				"column": 1,
				"source": map[string]interface{}{
					"path": frame.Position.File,
					"name": frame.Position.File,
				},
			},
		},
		"totalFrames": 1,
	}, "")
}

func (s *DAPSession) handleScopes(req dapMessage) {
	s.sendResponse(req, true, map[string]interface{}{
		"scopes": []map[string]interface{}{
			{"name": "Locals", "variablesReference": 1, "expensive": false},
			{"name": "Globals", "variablesReference": 2, "expensive": true},
		},
	}, "")
}

func (s *DAPSession) handleVariables(req dapMessage) {
	var args struct {
		VariablesReference int `json:"variablesReference"`
	}
	_ = json.Unmarshal(req.Arguments, &args)
	frame := s.debugger.LastFrame()
	if frame == nil {
		s.sendResponse(req, true, map[string]interface{}{"variables": []interface{}{}}, "")
		return
	}
	var source []Variable
	if args.VariablesReference == 2 {
		source = frame.Globals
	} else {
		source = frame.Locals
	}
	vars := make([]map[string]interface{}, 0, len(source))
	for _, variable := range source {
		vars = append(vars, map[string]interface{}{
			"name":               variable.Name,
			"type":               variable.Type,
			"value":              fmt.Sprint(variable.Value),
			"variablesReference": 0,
		})
	}
	s.sendResponse(req, true, map[string]interface{}{"variables": vars}, "")
}

func (s *DAPSession) readMessage() (dapMessage, error) {
	contentLength := -1
	for {
		line, err := s.in.ReadString('\n')
		if err != nil {
			return dapMessage{}, err
		}
		line = strings.TrimRight(line, "\r\n")
		if line == "" {
			break
		}
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}
		if strings.EqualFold(strings.TrimSpace(parts[0]), "Content-Length") {
			length, err := strconv.Atoi(strings.TrimSpace(parts[1]))
			if err != nil {
				return dapMessage{}, err
			}
			contentLength = length
		}
	}
	if contentLength < 0 {
		return dapMessage{}, fmt.Errorf("missing Content-Length header")
	}
	body := make([]byte, contentLength)
	if _, err := io.ReadFull(s.in, body); err != nil {
		return dapMessage{}, err
	}
	var msg dapMessage
	if err := json.Unmarshal(body, &msg); err != nil {
		return dapMessage{}, err
	}
	return msg, nil
}

func (s *DAPSession) sendResponse(req dapMessage, success bool, body interface{}, message string) {
	resp := dapMessage{
		Type:       "response",
		RequestSeq: req.Seq,
		Command:    req.Command,
		Success:    success,
		Body:       body,
		Message:    message,
	}
	s.send(resp)
}

func (s *DAPSession) sendEvent(event string, body interface{}) {
	s.send(dapMessage{
		Type:  "event",
		Event: event,
		Body:  body,
	})
}

func (s *DAPSession) send(msg dapMessage) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.seq++
	msg.Seq = s.seq
	payload, _ := json.Marshal(msg)
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "Content-Length: %d\r\n\r\n", len(payload))
	buf.Write(payload)
	_, _ = s.out.Write(buf.Bytes())
}

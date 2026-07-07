package websocket

import (
	"fmt"
	"sync"

	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"
	"github.com/gorilla/websocket"
)

// WebSocketModule 为 Lua VM 暴露真实 WebSocket 连接能力。
type WebSocketModule struct {
	connections map[int]*websocket.Conn
	nextHandle  int
	mu          sync.Mutex
	callbackMu  sync.Mutex
}

func New() *WebSocketModule { return &WebSocketModule{} }

func (m *WebSocketModule) Name() string { return "websocket" }

func (m *WebSocketModule) IsAvailable() bool { return true }

// Register 注册 websocket.connect/send/close，并保持旧绑定的对象生命周期语义。
func (m *WebSocketModule) Register(engine model.Engine) error {
	m.ensureState()

	engine.RegisterMethod("websocket.connect", "连接 WebSocket 服务器", m.connect, true)
	engine.RegisterMethod("websocket.connectObject", "连接 WebSocket 服务器并返回连接对象", m.connect, true)
	engine.RegisterMethod("websocket.close", "关闭 WebSocket 连接", m.close, true)
	engine.RegisterMethod("websocket.send", "发送 WebSocket 文本消息", m.send, true)
	return nil
}

func GetModule() model.Module { return &WebSocketModule{} }

func (m *WebSocketModule) ensureState() {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.connections == nil {
		m.connections = make(map[int]*websocket.Conn)
		m.nextHandle = 1
	}
}

func (m *WebSocketModule) connect(url string, onOpened func(map[string]interface{}), onClosed func(int), onError func(int, string), onRecv func(int, string)) (map[string]interface{}, error) {
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		if onError != nil {
			m.callError(onError, 0, err.Error())
		}
		return nil, err
	}

	handle := m.store(conn)
	connection := m.wrapConnection(handle)
	go m.readLoop(handle, conn, onClosed, onError, onRecv)

	if onOpened != nil {
		m.callbackMu.Lock()
		onOpened(connection)
		m.callbackMu.Unlock()
	}
	return connection, nil
}

func (m *WebSocketModule) store(conn *websocket.Conn) int {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.connections == nil {
		m.connections = make(map[int]*websocket.Conn)
		m.nextHandle = 1
	}
	handle := m.nextHandle
	m.nextHandle++
	m.connections[handle] = conn
	return handle
}

func (m *WebSocketModule) wrapConnection(handle int) map[string]interface{} {
	return map[string]interface{}{
		"handle": handle,
		"send": func(text string) bool {
			ok, err := m.send(handle, text)
			return err == nil && ok
		},
		"close": func() {
			_ = m.close(handle)
		},
	}
}

func (m *WebSocketModule) readLoop(handle int, conn *websocket.Conn, onClosed func(int), onError func(int, string), onRecv func(int, string)) {
	defer func() {
		_ = m.close(handle)
		if onClosed != nil {
			m.callbackMu.Lock()
			onClosed(handle)
			m.callbackMu.Unlock()
		}
	}()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			if onError != nil {
				m.callError(onError, handle, err.Error())
			}
			return
		}
		if onRecv != nil {
			m.callbackMu.Lock()
			onRecv(handle, string(message))
			m.callbackMu.Unlock()
		}
	}
}

func (m *WebSocketModule) callError(onError func(int, string), handle int, message string) {
	m.callbackMu.Lock()
	defer m.callbackMu.Unlock()
	onError(handle, message)
}

func (m *WebSocketModule) close(handle int) error {
	m.mu.Lock()
	conn, ok := m.connections[handle]
	if ok {
		delete(m.connections, handle)
	}
	m.mu.Unlock()

	if !ok {
		return nil
	}
	return conn.Close()
}

func (m *WebSocketModule) send(handle int, text string) (bool, error) {
	m.mu.Lock()
	conn, ok := m.connections[handle]
	m.mu.Unlock()
	if !ok {
		return false, fmt.Errorf("invalid websocket handle: %d", handle)
	}
	if err := conn.WriteMessage(websocket.TextMessage, []byte(text)); err != nil {
		return false, err
	}
	return true, nil
}

package websocket

import (
	"fmt"
	"sync"

	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"
	"github.com/gorilla/websocket"
	lua "github.com/yuin/gopher-lua"
)

// WebSocketModule websocket 模块
type WebSocketModule struct {
	connections map[int]*websocket.Conn
	nextHandle  int
	mu          sync.Mutex
}

// Name 返回模块名称
func (m *WebSocketModule) Name() string {
	return "websocket"
}

// IsAvailable 检查模块是否可用
func (m *WebSocketModule) IsAvailable() bool {
	return true
}

// Register 向引擎注册方法
func (m *WebSocketModule) Register(engine model.Engine) error {
	state := engine.GetState()

	m.mu.Lock()
	if m.connections == nil {
		m.connections = make(map[int]*websocket.Conn)
		m.nextHandle = 1
	}
	m.mu.Unlock()

	wsObj := state.NewTable()
	state.SetGlobal("websocket", wsObj)

	wsObj.RawSetString("connect", state.NewFunction(func(L *lua.LState) int {
		urlStr := L.CheckString(1)
		onOpened := L.CheckFunction(2)
		onClosed := L.CheckFunction(3)
		onError := L.CheckFunction(4)
		onRecv := L.CheckFunction(5)

		conn, _, err := websocket.DefaultDialer.Dial(urlStr, nil)
		if err != nil {
			if onError != nil {
				L.Push(onError)
				L.Push(lua.LNumber(0))
				L.Push(lua.LString(err.Error()))
				L.Call(2, 0)
			}
			L.Push(lua.LNumber(0))
			return 1
		}

		m.mu.Lock()
		handle := m.nextHandle
		m.nextHandle++
		m.connections[handle] = conn
		m.mu.Unlock()

		go func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Printf("panic in websocket read loop: %v\n", r)
				}
			}()

			for {
				_, message, err := conn.ReadMessage()
				if err != nil {
					if onError != nil {
						thread, _ := L.NewThread()
						thread.Push(onError)
						thread.Push(lua.LNumber(handle))
						thread.Push(lua.LString(err.Error()))
						if err := thread.PCall(2, 0, nil); err != nil {
							fmt.Printf("Error calling onError callback: %v\n", err)
						}
					}
					if onClosed != nil {
						thread, _ := L.NewThread()
						thread.Push(onClosed)
						thread.Push(lua.LNumber(handle))
						if err := thread.PCall(1, 0, nil); err != nil {
							fmt.Printf("Error calling onClosed callback: %v\n", err)
						}
					}
					break
				}

				if onRecv != nil {
					thread, _ := L.NewThread()
					thread.Push(onRecv)
					thread.Push(lua.LNumber(handle))
					thread.Push(lua.LString(string(message)))
					if err := thread.PCall(2, 0, nil); err != nil {
						fmt.Printf("Error calling onRecv callback: %v\n", err)
					}
				}
			}
		}()

		if onOpened != nil {
			L.Push(onOpened)
			L.Push(lua.LNumber(handle))
			if err := L.PCall(1, 0, nil); err != nil {
				fmt.Printf("Error calling onOpened callback: %v\n", err)
			}
		}

		L.Push(lua.LNumber(handle))
		return 1
	}))

	wsObj.RawSetString("close", state.NewFunction(func(L *lua.LState) int {
		handle := L.CheckInt(1)
		m.mu.Lock()
		defer m.mu.Unlock()

		if conn, ok := m.connections[handle]; ok {
			conn.Close()
			delete(m.connections, handle)
		}
		return 0
	}))

	wsObj.RawSetString("send", state.NewFunction(func(L *lua.LState) int {
		handle := L.CheckInt(1)
		text := L.CheckString(2)

		m.mu.Lock()
		conn, ok := m.connections[handle]
		m.mu.Unlock()

		if !ok {
			L.Push(lua.LBool(false))
			return 1
		}

		err := conn.WriteMessage(websocket.TextMessage, []byte(text))
		if err != nil {
			L.Push(lua.LBool(false))
			return 1
		}

		L.Push(lua.LBool(true))
		return 1
	}))

	engine.RegisterMethod("websocket.connect", "连接 WebSocket 服务器", func(url string, onOpened, onClosed, onError, onRecv lua.LValue) (int, error) {
		conn, _, err := websocket.DefaultDialer.Dial(url, nil)
		if err != nil {
			return 0, err
		}

		m.mu.Lock()
		handle := m.nextHandle
		m.nextHandle++
		m.connections[handle] = conn
		m.mu.Unlock()

		go func() {
			for {
				_, message, err := conn.ReadMessage()
				if err != nil {
					break
				}
				if onRecv != nil {
					if fn, ok := onRecv.(*lua.LFunction); ok {
						state := engine.GetState()
						thread, _ := state.NewThread()
						thread.Push(fn)
						thread.Push(lua.LNumber(handle))
						thread.Push(lua.LString(string(message)))
						if err := thread.PCall(2, 0, nil); err != nil {
							fmt.Printf("Error calling onRecv callback: %v\n", err)
						}
					}
				}
			}
		}()

		if onOpened != nil {
			if fn, ok := onOpened.(*lua.LFunction); ok {
				state := engine.GetState()
				state.Push(fn)
				state.Push(lua.LNumber(handle))
				if err := state.PCall(1, 0, nil); err != nil {
					fmt.Printf("Error calling onOpened callback: %v\n", err)
				}
			}
		}

		return handle, nil
	}, true)

	engine.RegisterMethod("websocket.close", "关闭 WebSocket 连接", func(handle int) error {
		m.mu.Lock()
		defer m.mu.Unlock()

		if conn, ok := m.connections[handle]; ok {
			conn.Close()
			delete(m.connections, handle)
		}
		return nil
	}, true)

	engine.RegisterMethod("websocket.send", "发送 WebSocket 消息", func(handle int, text string) (bool, error) {
		m.mu.Lock()
		conn, ok := m.connections[handle]
		m.mu.Unlock()

		if !ok {
			return false, fmt.Errorf("invalid handle")
		}

		err := conn.WriteMessage(websocket.TextMessage, []byte(text))
		if err != nil {
			return false, err
		}

		return true, nil
	}, true)

	return nil
}

// GetModule 获取模块实例
func GetModule() model.Module {
	return New()
}

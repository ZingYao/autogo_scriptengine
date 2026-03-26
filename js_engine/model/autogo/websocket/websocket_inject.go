package websocket

import (
	"fmt"
	"sync"

	"github.com/ZingYao/autogo_scriptengine/js_engine/model"
	"github.com/dop251/goja"
	"github.com/gorilla/websocket"
)

// WebSocketModule websocket 模块
type WebSocketModule struct {
	connections map[int]*websocket.Conn
	nextHandle  int
	mu          sync.Mutex
}

// New 创建新的 WebSocket 模块实例
func New() *WebSocketModule {
	return &WebSocketModule{
		connections: make(map[int]*websocket.Conn),
		nextHandle:  1,
	}
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
	vm := engine.GetVM()

	wsObj := vm.NewObject()
	vm.Set("websocket", wsObj)

	wsObj.Set("connect", func(call goja.FunctionCall) goja.Value {
		url := call.Argument(0).String()
		onOpened := call.Argument(1)
		onClosed := call.Argument(2)
		onError := call.Argument(3)
		onRecv := call.Argument(4)

		conn, _, err := websocket.DefaultDialer.Dial(url, nil)
		if err != nil {
			if onError != nil && !goja.IsUndefined(onError) {
				if fn, ok := goja.AssertFunction(onError); ok {
					fn(nil, vm.ToValue(0), vm.ToValue(err.Error()))
				}
			}
			return vm.ToValue(0)
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
					if onError != nil && !goja.IsUndefined(onError) {
						if fn, ok := goja.AssertFunction(onError); ok {
							fn(nil, vm.ToValue(handle), vm.ToValue(err.Error()))
						}
					}
					if onClosed != nil && !goja.IsUndefined(onClosed) {
						if fn, ok := goja.AssertFunction(onClosed); ok {
							fn(nil, vm.ToValue(handle))
						}
					}
					break
				}

				if onRecv != nil && !goja.IsUndefined(onRecv) {
					if fn, ok := goja.AssertFunction(onRecv); ok {
						fn(nil, vm.ToValue(handle), vm.ToValue(string(message)))
					}
				}
			}
		}()

		if onOpened != nil && !goja.IsUndefined(onOpened) {
			if fn, ok := goja.AssertFunction(onOpened); ok {
				fn(nil, vm.ToValue(handle))
			}
		}

		return vm.ToValue(handle)
	})

	wsObj.Set("close", func(call goja.FunctionCall) goja.Value {
		handle := int(call.Argument(0).ToInteger())
		m.mu.Lock()
		defer m.mu.Unlock()

		if conn, ok := m.connections[handle]; ok {
			conn.Close()
			delete(m.connections, handle)
		}
		return goja.Undefined()
	})

	wsObj.Set("send", func(call goja.FunctionCall) goja.Value {
		handle := int(call.Argument(0).ToInteger())
		text := call.Argument(1).String()

		m.mu.Lock()
		conn, ok := m.connections[handle]
		m.mu.Unlock()

		if !ok {
			return vm.ToValue(false)
		}

		err := conn.WriteMessage(websocket.TextMessage, []byte(text))
		if err != nil {
			return vm.ToValue(false)
		}

		return vm.ToValue(true)
	})

	engine.RegisterMethod("websocket.connect", "连接 WebSocket 服务器", func(url string, onOpened, onClosed, onError, onRecv goja.Value) (int, error) {
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
				if onRecv != nil && !goja.IsUndefined(onRecv) {
					if fn, ok := goja.AssertFunction(onRecv); ok {
						fn(nil, vm.ToValue(handle), vm.ToValue(string(message)))
					}
				}
			}
		}()

		if onOpened != nil && !goja.IsUndefined(onOpened) {
			if fn, ok := goja.AssertFunction(onOpened); ok {
				fn(nil, vm.ToValue(handle))
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

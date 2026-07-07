package http

import (
	"bytes"
	"io"
	nethttp "net/http"
	"time"

	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"
)

// HttpModule 是 go-lua-vm 迁移后的模块壳。
type HttpModule struct{}

func New() *HttpModule { return &HttpModule{} }

func (m *HttpModule) Name() string { return "https" }

func (m *HttpModule) IsAvailable() bool { return true }

func (m *HttpModule) Register(engine model.Engine) error {
	engine.RegisterMethod("https.get", "发送 GET 请求", func(url string, timeout ...int) map[string]interface{} {
		return request(nethttp.MethodGet, url, nil, nil, timeoutValue(timeout...))
	}, true)
	engine.RegisterMethod("https.post", "发送 POST 请求", func(url string, data string, headers map[string]string, timeout ...int) map[string]interface{} {
		return request(nethttp.MethodPost, url, []byte(data), headers, timeoutValue(timeout...))
	}, true)
	engine.RegisterMethod("https.postMultipart", "发送 Multipart POST 请求", func(url string, fileName string, fileData string, timeout ...int) map[string]interface{} {
		return request(nethttp.MethodPost, url, []byte(fileData), map[string]string{"Content-Type": "application/octet-stream"}, timeoutValue(timeout...))
	}, true)
	return nil
}

func GetModule() model.Module { return &HttpModule{} }

func timeoutValue(timeout ...int) int {
	if len(timeout) == 0 || timeout[0] <= 0 {
		return 5000
	}
	return timeout[0]
}

func request(method string, url string, data []byte, headers map[string]string, timeout int) map[string]interface{} {
	client := &nethttp.Client{Timeout: time.Duration(timeout) * time.Millisecond}
	req, err := nethttp.NewRequest(method, url, bytes.NewReader(data))
	if err != nil {
		return map[string]interface{}{"code": 0, "data": err.Error()}
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	resp, err := client.Do(req)
	if err != nil {
		return map[string]interface{}{"code": 0, "data": err.Error()}
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return map[string]interface{}{"code": resp.StatusCode, "data": err.Error()}
	}
	return map[string]interface{}{"code": resp.StatusCode, "data": string(body)}
}

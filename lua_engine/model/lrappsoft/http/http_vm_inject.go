package http

import (
	"io"
	nethttp "net/http"
	"strings"
	"time"

	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"
)

// HttpModule 是 go-lua-vm 迁移后的模块壳。
type HttpModule struct{}

func New() *HttpModule { return &HttpModule{} }

func (m *HttpModule) Name() string { return "http" }

func (m *HttpModule) IsAvailable() bool { return true }

func (m *HttpModule) Register(engine model.Engine) error {
	request := func(url, method string, headers map[string]string, body string, timeoutSeconds int, followRedirect bool) (int, string, map[string]string, string, error) {
		return httpRequest(url, method, headers, body, timeoutSeconds, followRedirect)
	}
	engine.RegisterMethod("http.request", "发送 HTTP 请求", request, true)
	engine.RegisterMethod("https.request", "发送 HTTPS 请求", request, true)
	engine.RegisterMethod("ssl.https.request", "发送 HTTPS 请求", request, true)
	return nil
}

func GetModule() model.Module { return &HttpModule{} }

func httpRequest(url, method string, headers map[string]string, body string, timeoutSeconds int, followRedirect bool) (int, string, map[string]string, string, error) {
	if method == "" {
		method = nethttp.MethodGet
	}
	if timeoutSeconds <= 0 {
		timeoutSeconds = 30
	}
	client := &nethttp.Client{Timeout: time.Duration(timeoutSeconds) * time.Second}
	if !followRedirect {
		client.CheckRedirect = func(_ *nethttp.Request, _ []*nethttp.Request) error {
			return nethttp.ErrUseLastResponse
		}
	}
	request, err := nethttp.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		return 0, "", nil, "", err
	}
	for key, value := range headers {
		request.Header.Set(key, value)
	}
	response, err := client.Do(request)
	if err != nil {
		return 0, "", nil, "", err
	}
	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return 0, "", nil, "", err
	}
	responseHeaders := make(map[string]string, len(response.Header))
	for key, values := range response.Header {
		if len(values) > 0 {
			responseHeaders[key] = values[0]
		}
	}
	return response.StatusCode, response.Status, responseHeaders, string(responseBody), nil
}

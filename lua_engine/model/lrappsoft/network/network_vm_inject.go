package network

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"
)

// NetworkModule 是 go-lua-vm 迁移后的模块壳。
type NetworkModule struct{}

func New() *NetworkModule { return &NetworkModule{} }

func (m *NetworkModule) Name() string { return "network" }

func (m *NetworkModule) IsAvailable() bool { return true }

func (m *NetworkModule) Register(engine model.Engine) error {
	engine.RegisterMethod("network.httpGet", "发送 HTTP GET 请求", func(url string, timeoutSeconds ...int) (map[string]interface{}, error) {
		return networkRequest(http.MethodGet, url, "", "", networkTimeout(timeoutSeconds...))
	}, true)
	engine.RegisterMethod("network.httpPost", "发送 HTTP POST 请求", func(url, body string, timeoutSeconds ...int) (map[string]interface{}, error) {
		return networkRequest(http.MethodPost, url, body, "application/x-www-form-urlencoded", networkTimeout(timeoutSeconds...))
	}, true)
	engine.RegisterMethod("network.httpPostData", "发送带 Content-Type 的 HTTP POST 请求", func(url, body, contentType string) (string, error) {
		response, err := networkRequest(http.MethodPost, url, body, contentType, 30)
		if err != nil {
			return "", err
		}
		return response["body"].(string), nil
	}, true)
	engine.RegisterMethod("network.downloadFile", "下载文件", func(url, path string) (bool, error) {
		return true, networkDownload(url, path)
	}, true)
	engine.RegisterMethod("network.httpDownload", "下载文件", func(url, path string) (bool, error) {
		return true, networkDownload(url, path)
	}, true)
	engine.RegisterMethod("network.uploadFile", "上传文件", func(url, path string) (string, error) {
		body, _, err := networkUpload(url, "file", path, nil)
		return body, err
	}, true)
	engine.RegisterMethod("network.httpUpload", "上传文件并返回响应体与状态码", func(url, fieldName, path string, headers map[string]string) (string, int, error) {
		return networkUpload(url, fieldName, path, headers)
	}, true)
	return nil
}

func GetModule() model.Module { return &NetworkModule{} }

func networkTimeout(timeoutSeconds ...int) int {
	if len(timeoutSeconds) == 0 || timeoutSeconds[0] <= 0 {
		return 30
	}
	return timeoutSeconds[0]
}

func networkRequest(method, url, body, contentType string, timeoutSeconds int) (map[string]interface{}, error) {
	client := &http.Client{Timeout: time.Duration(timeoutSeconds) * time.Second}
	var reader io.Reader
	if body != "" {
		reader = strings.NewReader(body)
	}
	request, err := http.NewRequest(method, url, reader)
	if err != nil {
		return nil, err
	}
	if contentType != "" {
		request.Header.Set("Content-Type", contentType)
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{
		"body":       string(responseBody),
		"statusCode": response.StatusCode,
		"status":     response.Status,
	}, nil
}

func networkDownload(url, path string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o644)
}

func networkUpload(url, fieldName, path string, headers map[string]string) (string, int, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	file, err := os.Open(path)
	if err != nil {
		return "", 0, err
	}
	defer file.Close()
	part, err := writer.CreateFormFile(fieldName, filepath.Base(path))
	if err != nil {
		return "", 0, err
	}
	if _, err := io.Copy(part, file); err != nil {
		return "", 0, err
	}
	if err := writer.Close(); err != nil {
		return "", 0, err
	}
	request, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return "", 0, err
	}
	request.Header.Set("Content-Type", writer.FormDataContentType())
	for key, value := range headers {
		request.Header.Set(key, value)
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return "", 0, err
	}
	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return "", response.StatusCode, err
	}
	return string(responseBody), response.StatusCode, nil
}

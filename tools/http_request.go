package tools

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

// HTTPRequest 发起http请求的接口
type HTTPRequest interface {
	Get() ([]byte, error)
	Post() ([]byte, error)
}

type httpRequest struct {
	url  string
	data string
}

// NewRequest 初始化HTTPRequest接口
func NewRequest(url string, data string) HTTPRequest {
	return &httpRequest{
		url:  url,
		data: data,
	}
}

//HTTPGet 不想写注释，破工具老是标黄线。
func (h *httpRequest) Get() ([]byte, error) {
	resp, err := http.Get(h.url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func (h *httpRequest) Post() ([]byte, error) {
	resquest, err := http.NewRequest("POST", h.url, bytes.NewReader([]byte(h.data)))
	if err != nil {
		panic(err)
	}
	resquest.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := &http.Client{}
	resp, err := client.Do(resquest)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

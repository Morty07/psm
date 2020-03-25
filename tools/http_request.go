package tools

import (
	"io/ioutil"
	"net/http"
)

//HTTPGet 不想写注释，破工具老是标黄线。
func HTTPGet(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

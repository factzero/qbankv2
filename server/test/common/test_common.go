package common

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func HttpPost(url string, data interface{}, token string) string {
	jsonStr, _ := json.Marshal(data)

	// http.Clicent:是一个HTTP客户端，客户机比往返器(比如传输)更高级,还处理HTTP细节，比如cookie，重定向，长连接。
	client := &http.Client{}
	// 使用 NewRequest 设置头参数、cookie之类的数据，
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	//设置请求头header,携带token身份认证
	req.Header.Set("authorization", token)

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return string(body)
}

func HttpGet(url string, token string) string {
	// http.Clicent:是一个HTTP客户端，客户机比往返器(比如传输)更高级,还处理HTTP细节，比如cookie，重定向，长连接。
	client := &http.Client{}
	// 使用 NewRequest 设置头参数、cookie之类的数据，
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	//设置请求头header,携带token身份认证
	req.Header.Set("authorization", token)

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return string(body)
}

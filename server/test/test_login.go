package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type loginResp struct {
	Code    int       `json:"code"`
	Data    string    `json:"data"`
	Exdata  any       `json:"exdata"`
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
	Token   any       `json:"token"`
}

func httpPost(url string, data interface{}) string {
	jsonStr, _ := json.Marshal(data)

	// http.Clicent:是一个HTTP客户端，客户机比往返器(比如传输)更高级,还处理HTTP细节，比如cookie，重定向，长连接。
	client := &http.Client{}
	// 使用 NewRequest 设置头参数、cookie之类的数据，
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return string(body)
}

func httpGet(url string, token string) string {
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

func main() {
	data := map[string]string{
		"username": "admin",
		"password": "admin",
	}

	res3 := httpPost("http://127.0.0.1:9000/admin/user/login", data)
	fmt.Println(res3)
	var res3_json loginResp
	err := json.Unmarshal([]byte(res3), &res3_json)
	if err != nil {
		fmt.Println("errrr")
	}
	token := res3_json.Data
	fmt.Println(token)

	res4 := httpGet("http://127.0.0.1:9000/admin/user/get_userinfo", token)
	fmt.Println(res4)
}

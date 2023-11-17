package common

import (
	"encoding/json"
	"fmt"
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

func TestLogin() string {
	data := map[string]string{
		"username": "admin",
		"password": "admin",
	}

	res := HttpPost("http://127.0.0.1:9000/admin/user/login", data, "")
	fmt.Println(res)
	var res_json loginResp
	err := json.Unmarshal([]byte(res), &res_json)
	if err != nil {
		fmt.Println("errrr")
	}
	token := res_json.Data

	// res2 := HttpGet("http://127.0.0.1:9000/admin/user/get_userinfo", token)
	// fmt.Println(res2)

	return token
}

package common

import (
	"fmt"
)

// 选择题
type ChoiceQuestion struct {
	Index   string `json:"index"`   // 序号
	Content string `json:"content"` // 选项内容
}

// 试题信息
type QuestionData struct {
	Qtype    string           `json:"qtype"`            // 题型
	Stem     string           `json:"stem"`             // 题干
	Answer   string           `json:"answer"`           // 答案
	Analysis string           `json:"analysis"`         // 解析
	Option   []ChoiceQuestion `json:"option,omitempty"` // 选择题
}

func TestManualEntryQues(token string) {
	data := QuestionData{
		Qtype:    "blank",
		Stem:     "ssdfadf",
		Answer:   "asdfadf",
		Analysis: "dh6776",
		Option:   []ChoiceQuestion{ChoiceQuestion{"A", "ssdfad"}, ChoiceQuestion{"B", "dsfa"}},
	}

	res := HttpPost("http://127.0.0.1:9000/admin/managequestion/managequestion/manualEntry", data, token)
	fmt.Println(res)
}

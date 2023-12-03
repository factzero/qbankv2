package managequestion

import (
	"encoding/json"
	"fmt"
	"gozero/dbmodel"
	"gozero/utils"
	"gozero/utils/results"
	"io"
	"reflect"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

/**
* 使用 Index 是省略路径中的index
* 本路径为： /admin/user/login -省去了index
**/
func init() {
	utils.Register(&ManageQuestion{}, reflect.TypeOf(ManageQuestion{}).PkgPath())
}

type ManageQuestion struct {
}

// 选择题
type choiceQuestion struct {
	Index   string `json:"index"`   // 序号
	Content string `json:"content"` // 选项内容
}

// 试题信息
type questionData struct {
	Qtype    string           `json:"qtype"`            // 题型:radio,multiple,determine,blanks,analytic
	Stem     string           `json:"stem"`             // 题干
	Answer   string           `json:"answer"`           // 答案
	Analysis string           `json:"analysis"`         // 解析
	Option   []choiceQuestion `json:"option,omitempty"` // 选择题
}

type pageInfo struct {
	Page     int    `json:"page"`     // 页码
	PageSize int    `json:"pageSize"` // 页码大小
	Keyword  string `json:"keyword"`  // 关键字
}

/**
* 手动录入试题
**/
func (api *ManageQuestion) ManualEntry(c *gin.Context) {
	// 获取post传过来的data
	body, _ := io.ReadAll(c.Request.Body)
	var parameter questionData
	_ = json.Unmarshal(body, &parameter)
	fmt.Println(parameter)
	option_cnt := len(parameter.Option)
	options := [8]string{}
	for i := 0; i < option_cnt; i++ {
		options[i] = parameter.Option[i].Content
	}
	if parameter.Qtype == "" || parameter.Stem == "" || parameter.Answer == "" {
		results.Failed(c, "试题录入失败：题干或者答案为空", nil)
		return
	}
	optionAnsList := "ABCDEFGH"
	switch parameter.Qtype {
	case "radio":
		if len(parameter.Answer) != 1 {
			results.Failed(c, "试题录入失败：单选题答案数量不正确", nil)
			return
		}
		if option_cnt <= 1 {
			results.Failed(c, "试题录入失败：单选题选项数量不正确", nil)
			return
		}
		parameter.Answer = strings.ToUpper(parameter.Answer)
		if strings.Count(optionAnsList[0:option_cnt], parameter.Answer) == 0 {
			results.Failed(c, "试题录入失败：单选题答案不在选项中", nil)
			return
		}
	case "multiple":
		if option_cnt < len(parameter.Answer) {
			results.Failed(c, "试题录入失败：多选题选项数量不正确", nil)
			return
		}
		parameter.Answer = strings.ToUpper(parameter.Answer)
		for _, val := range parameter.Answer {
			if strings.Count(optionAnsList[0:option_cnt], string(val)) == 0 {
				results.Failed(c, "试题录入失败：多选题答案不在选项中", nil)
				return
			}
		}
	case "determine", "blanks", "analytic":
		for i := 0; i < 8; i++ {
			options[i] = ""
		}
		option_cnt = 0
	default:
		results.Failed(c, "试题录入失败：题型错误", nil)
		return
	}
	res, err := dbmodel.DB().Table("question_data").
		Data(map[string]interface{}{"qtype": parameter.Qtype, "stem": parameter.Stem, "answer": parameter.Answer,
			"analysis": parameter.Analysis, "option_cnt": option_cnt, "option_A": options[0], "option_B": options[1],
			"option_C": options[2], "option_D": options[3], "option_E": options[4], "option_F": options[5],
			"option_G": options[6], "option_H": options[7], "createtime": time.Now()}).Insert()
	if err != nil || res == 0 {
		results.Failed(c, "试题录入失败：写入写入数据库失败", nil)
		return
	}

	results.Success(c, "试题录入成功", nil, nil)
}

/**
* 查询试题
**/
func (api *ManageQuestion) Get_questions(c *gin.Context) {
	// 获取post传过来的data
	body, _ := io.ReadAll(c.Request.Body)
	var parameter pageInfo
	_ = json.Unmarshal(body, &parameter)
	fmt.Println(parameter)
	pageNo := parameter.Page
	pageSize := parameter.PageSize
	MDB := dbmodel.DB().Table("question_data").Fields("id, qtype, stem, createtime, updatetime")
	MDBC := dbmodel.DB().Table("question_data")
	list, err := MDB.Limit(pageSize).Page(pageNo).Get()
	if err != nil {
		results.Failed(c, err.Error(), nil)
	}
	totalCount, _ := MDBC.Count("*")

	results.Success(c, "试题查询成功！", map[string]interface{}{
		"page":     pageNo,
		"pageSize": pageSize,
		"total":    totalCount,
		"item":     list}, nil)
}

package managequestion

import (
	"encoding/json"
	"fmt"
	"gozero/dbmodel"
	"gozero/utils"
	"gozero/utils/results"
	"io"
	"reflect"
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
	Qtype    string           `json:"qtype"`            // 题型
	Stem     string           `json:"stem"`             // 题干
	Answer   string           `json:"answer"`           // 答案
	Analysis string           `json:"analysis"`         // 解析
	Option   []choiceQuestion `json:"option,omitempty"` // 选择题
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
	dbmodel.DB().Table("question_data").
		Data(map[string]interface{}{"qtype": parameter.Qtype, "stem": parameter.Stem, "answer": parameter.Answer,
			"analysis": parameter.Analysis, "option_cnt": option_cnt, "option_A": options[0], "option_B": options[1],
			"option_C": options[2], "option_D": options[3], "option_E": options[4], "option_F": options[5],
			"option_G": options[6], "option_H": options[7], "createtime": time.Now()}).Insert()

	results.Success(c, "试题录入成功返回！", nil, nil)
}

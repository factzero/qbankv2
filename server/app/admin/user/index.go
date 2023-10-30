package user

import (
	"encoding/json"
	"gozero/route/middleware"
	"gozero/utils"
	"gozero/utils/results"
	"io"
	"reflect"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

/**
* 使用 Index 是省略路径中的index
* 本路径为： /admin/user/login -省去了index
**/
func init() {
	utils.Register(&Index{}, reflect.TypeOf(Index{}).PkgPath())
}

type Index struct {
}

/**
* 登录
**/
func (api *Index) Login(c *gin.Context) {
	//获取post传过来的data
	body, _ := io.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	if parameter["username"] == nil || parameter["password"] == nil {
		results.Failed(c, "请提交用户账号或密码！", nil)
		return
	}
	// token
	token := middleware.GenerateToken(&middleware.UserClaims{
		ID:             13413413,
		StandardClaims: jwt.StandardClaims{},
	})

	results.Success(c, "登录成功返回token！", token, nil)
}

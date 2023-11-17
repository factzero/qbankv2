package user

import (
	"encoding/json"
	"gozero/dbmodel"
	"gozero/route/middleware"
	"gozero/utils"
	"gozero/utils/results"
	"io"
	"reflect"
	"time"

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
	// 获取post传过来的data
	body, _ := io.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	if parameter["username"] == nil || parameter["password"] == nil {
		results.Failed(c, "请提交用户账号或密码！", nil)
		return
	}
	username := parameter["username"].(string)
	password := parameter["password"].(string)
	res, err := dbmodel.DB().Table("admin_user").Fields("id,password").Where("username", username).OrWhere("email", username).First()
	if res == nil || err != nil {
		results.Failed(c, "账号不存在！", nil)
		return
	}
	if password != res["password"] {
		results.Failed(c, "您输入的密码不正确！", password)
		return
	}
	// token
	token := middleware.GenerateToken(&middleware.UserClaims{
		ID:             res["id"].(int64),
		StandardClaims: jwt.StandardClaims{},
	})

	// 登录日志
	dbmodel.DB().Table("admin_user").Where("id", res["id"]).
		Data(map[string]interface{}{"lastLoginTime": time.Now(), "lastLoginIp": utils.GetRealIP(c)}).Update()
	dbmodel.DB().Table("login_logs").
		Data(map[string]interface{}{"uid": res["id"], "out_in": "in",
			"createtime": time.Now(), "loginIP": utils.GetRealIP(c)}).Insert()

	results.Success(c, "登录成功返回token！", token, nil)
}

/**
* 获取用户信息
**/
func (api *Index) Get_userinfo(c *gin.Context) {
	getuser, _ := c.Get("user")
	user := getuser.(*middleware.UserClaims)
	userdata, err := dbmodel.DB().Table("admin_user").Fields("id,username,email").Where("id", user.ID).First()
	if err != nil {
		results.Failed(c, "获取用户信息失败!", err)
	} else {
		results.Success(c, "获取用户信息成功!", map[string]interface{}{
			"userID":   userdata["id"],
			"username": userdata["username"],
		}, nil)
	}
}

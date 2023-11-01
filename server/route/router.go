package route

import (
	_ "gozero/app"
	"gozero/route/middleware"
	"gozero/utils"

	"github.com/gin-gonic/gin"
)

// 路由初始化
func InitRouter() *gin.Engine {
	// 初始化路由
	R := gin.Default()

	// 验证token
	R.Use(middleware.JwtVerify)

	// 绑定自定义路由
	utils.Bind(R)
	return R
}

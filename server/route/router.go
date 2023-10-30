package route

import (
	_ "gozero/app"
	"gozero/utils"

	"github.com/gin-gonic/gin"
)

// 路由初始化
func InitRouter() *gin.Engine {
	// 初始化路由
	R := gin.Default()
	// 绑定自定义路由
	utils.Bind(R)
	return R
}

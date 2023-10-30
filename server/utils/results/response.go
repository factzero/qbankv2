package results

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 请求成功,使用该方法返回信息
func Success(ctx *gin.Context, msg string, data interface{}, exdata interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": msg,
		"data":    data,
		"exdata":  exdata,
		"token":   "newtoken",
		"time":    time.Now().Unix(),
	})
}

// 请求失败, 使用该方法返回信息
func Failed(ctx *gin.Context, msg string, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": msg,
		"data":    data,
		"time":    time.Now().Unix(),
	})
}

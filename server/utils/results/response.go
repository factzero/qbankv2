package results

import (
	"gozero/route/middleware"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 请求成功,使用该方法返回信息
func Success(c *gin.Context, msg string, data interface{}, exdata interface{}) {
	token := c.Request.Header.Get("Authorization")
	var newtoken interface{}
	if token != "" {
		tokentmp := middleware.RefreshToken(token)
		if tokentmp != nil {
			newtoken = tokentmp
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": msg,
		"data":    data,
		"exdata":  exdata,
		"token":   newtoken,
		"time":    time.Now(),
	})
}

// 请求失败, 使用该方法返回信息
func Failed(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": msg,
		"data":    data,
		"time":    time.Now(),
	})
}

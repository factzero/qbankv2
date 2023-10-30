package middleware

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// 用户信息类，作为生成token的参数
type UserClaims struct {
	ID                 int64  `json:"id"`
	Username           string `json:"username"`
	jwt.StandardClaims        //jwt-go提供的标准claim
}

var secret = []byte("45245246566ghdfghd") // 自定义的token秘钥
var effectTime = 120 * time.Minute        // 分钟单位

// 生成token
func GenerateToken(claims *UserClaims) interface{} {
	// 设置token有效期，也可不设置有效期，采用redis的方式
	//   1)将token存储在redis中，设置过期时间，token如没过期，则自动刷新redis过期时间，
	//   2)通过这种方式，可以很方便的为token续期，而且也可以实现长时间不登录的话，强制登录
	// 本例只是简单采用 设置token有效期的方式，只是提供了刷新token的方法，并没有做续期处理的逻辑
	claims.ExpiresAt = time.Now().Add(effectTime).Unix()
	// 生成token
	sign, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secret)
	if err != nil {
		//这里因为项目接入了统一异常处理，所以使用panic并不会使程序终止，如不接入，可使用原始方式处理错误
		//接入统一异常可参考 https://blog.csdn.net/u014155085/article/details/106733391
		panic(err)
	}
	return sign
}

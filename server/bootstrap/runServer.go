package bootstrap

import (
	"gozero/dbmodel"
	"gozero/route"
)

func RunServer() {
	//加载路由
	r := route.InitRouter()
	dbmodel.DBInit(1)
	r.Run(":9000")
}

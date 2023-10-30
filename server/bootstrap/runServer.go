package bootstrap

import "gozero/route"

func RunServer() {
	//加载路由
	r := route.InitRouter()
	r.Run(":9000")
}

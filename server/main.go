package main

import (
	"gozero/bootstrap"
	"gozero/global"
)

func main() {
	// 初始化配置
	global.App.Config.InitializeConfig()
	// 初始化日志
	global.App.Log = bootstrap.InitializeLog()
	global.App.Log.Info("项目启动成功")
	// 启动服务器
	bootstrap.RunServer()
}

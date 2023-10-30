package main

import (
	"gozero/bootstrap"
	"gozero/global"
)

func main() {
	// 初始化配置
	global.App.Config.InitializeConfig()
	// 启动服务器
	bootstrap.RunServer()
}

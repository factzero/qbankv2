package dbmodel

import (
	"fmt"
	"gozero/global"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gohouse/gorose/v2"
)

var err error
var engin *gorose.Engin

// 取得数据库连接实例
func DBInit(starType interface{}) {
	global.App.Log.Info(fmt.Sprintf("连接数据库中:%v", starType))
	dsbSource := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local&timeout=1000ms", global.App.Config.DBconf.Username, global.App.Config.DBconf.Password, global.App.Config.DBconf.Hostname, global.App.Config.DBconf.Hostport, global.App.Config.DBconf.Database)
	engin, err = gorose.Open(&gorose.Config{Driver: global.App.Config.DBconf.Driver, Dsn: dsbSource, Prefix: global.App.Config.DBconf.Prefix})
	if err != nil {
		global.App.Log.Info(fmt.Sprintf("数据库连接实例错误: %v", err))
	} else {
		global.App.Log.Info(fmt.Sprintf("连接数据库成功:%v", starType))
		engin.GetExecuteDB().SetMaxIdleConns(10)                  // 连接池最大空闲连接数,不设置, 默认无
		engin.GetExecuteDB().SetMaxOpenConns(50)                  // 连接池最大连接数,不设置, 默认无限
		engin.GetExecuteDB().SetConnMaxLifetime(59 * time.Second) // 时间比超时时间短
		engin.GetQueryDB().Exec("SET @@sql_mode='NO_ENGINE_SUBSTITUTION';")
	}
}

// controller层调用
func DB() gorose.IOrm {
	return engin.NewOrm()
}

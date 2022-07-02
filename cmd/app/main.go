package main

import (
	"iris-starter/internal/config"
	"iris-starter/pkg/logging"

	"github.com/kataras/iris/v12"
)

var (
	err error
	app = iris.New()
)

func main() {
	defer func() {
		logging.Sync()
	}()

	// 初始化路由
	// api.Index(app)

	// 监听端口
	logging.Info("Start Web Server")
	if err = app.Run(iris.Addr(":"+config.GetConfig().HttpPort), iris.WithoutInterruptHandler); err != nil {
		logging.Error("Start Web Server err: " + err.Error())
	}
}

func init() {

	// 初始化日志
	logging.Init(config.GetConfig().Log)

	// 注册mysql

	// 注册redis

}

package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"

	"iris-starter/internal/config"
	"iris-starter/pkg/logging"

	"iris-starter/internal/router"
)

var (
	err error
	app = iris.New()
)

func main() {
	defer func() {
		logging.Sync()
	}()

	// 监听端口
	logging.Info("Start Web Server !")
	if err = app.Run(iris.Addr(":"+config.GetConfig().HttpPort), iris.WithoutInterruptHandler); err != nil {
		logging.Error("Start Web Server err: " + err.Error())
	}
}

func init() {
	// 引入参数校验
	app.Validator = validator.New()

	// 初始化日志
	logging.Init(config.GetConfig().Log)

	// 初始化路由
	router.Init(app)

	// 注册mysql

	// 注册redis

}

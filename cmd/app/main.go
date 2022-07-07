package main

import (
	"context"
	"time"

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

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /v2
func main() {
	defer func() {
		logging.Sync()
	}()

	// 优雅关闭服务
	iris.RegisterOnInterrupt(func() {
		timeout := 5 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		//关闭所有主机
		app.Shutdown(ctx)
	})

	// 监听端口
	logging.Info("Start Web Server !")
	if err = app.Run(iris.Addr(":"+config.GetConfig().HttpPort),
		iris.WithoutInterruptHandler,
		iris.WithOptimizations,
		iris.WithProtoJSON,
		iris.WithTimeout(time.Duration(config.GetConfig().HttpRequestTimeout)*time.Second)); err != nil {
		logging.Error("Start Server err: " + err.Error())
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

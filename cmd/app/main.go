package main

import (
	"context"
	"fmt"
	"os/exec"
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

// @title API DOC
// @version 1.0
// @description API文档.

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /v1
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
		iris.WithTimeFormat("2022-02-20 12:12:12"),
		iris.WithTimeout(time.Duration(config.GetConfig().HttpRequestTimeout)*time.Second)); err != nil {
		logging.Error("Start Server err: " + err.Error())
	}
}

func init() {

	// 初始化接口文档
	InitDoc()

	// 引入参数校验
	app.Validator = validator.New()

	// 初始化日志
	logging.Init(config.GetConfig().Log)

	// 初始化路由
	router.Init(app)

	// 注册mysql

	// 注册redis

}

func InitDoc() {
	_, err := exec.Command("./scripts/docs.sh").Output()
	if err != nil {
		fmt.Printf("error %s", err)
	}
}

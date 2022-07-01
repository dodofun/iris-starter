package main

import (
	"context"
	"iris-starter/internal/config"
	"iris-starter/pkg/logging"

	"sync"
	"time"

	"github.com/kataras/iris/v12"
)

var (
	err         error
	ctx, cancel = context.WithCancel(context.Background())
	wg          = new(sync.WaitGroup)
	app         = iris.New()
)

func main() {
	defer func() {
		wg.Wait()
	}()

	// 初始化路由
	// api.Index(app)

	// 监听端口
	logging.Info("Start Web Server ")
	if err = app.Run(iris.Addr(":"+config.GetConfig().HttpPort), iris.WithoutInterruptHandler); err != nil {
		logging.Error("Start Web Server err: " + err.Error())
	}
}

func init() {

	// 初始化日志
	logging.Init(config.GetConfig().Log)

	// 注册mysql

	// 注册redis

	// 优雅关闭程序
	iris.RegisterOnInterrupt(func() {
		wg.Add(1)
		defer wg.Done()
		cancel()
		time.Sleep(5 * time.Second)
		_ = app.Shutdown(ctx)
	})
}

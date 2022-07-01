package main

import (
	"context"
	"flag"
	"fmt"
	"iris-starter/internal/config"

	"sync"
	"time"

	"github.com/kataras/iris/v12"
)

var (
	err         error
	configPath  string
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
	// logging.Info("Start Web Server ")
	fmt.Println(config.GetConfig().AppName)
	if err = app.Run(iris.Addr(":"+config.GetConfig().HttpPort), iris.WithoutInterruptHandler); err != nil {
		fmt.Printf("Start Web Server err: " + err.Error())
	}
}

func init() {
	flag.StringVar(&configPath, "config", "./configs/config.json", "配置文件路径以及文件名(必填)")
	flag.Parse()

	// 初始化配置
	// config.Init(configPath)

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

package router

import (
	"iris-starter/internal/middleware"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/pprof"
)

func Init(app *iris.Application) {
	// 重复路由报错
	app.SetRegisterRule(iris.RouteError)

	// 引入中间件
	middleware.Init(app)

	// 服务健康检查
	app.Get("/health", func(ctx iris.Context) {
		ctx.StatusCode(iris.StatusOK)
	})

	// 性能分析
	p := pprof.New()
	app.Any("/debug/pprof", p)
	app.Any("/debug/pprof/{action:path}", p)

	// 注册服务资源
	registerRouter(app)

	// 注册静态资源
	app.HandleDir("/assets", iris.Dir("./views/assets"), iris.DirOptions{IndexName: "/index.html", Compress: true, SPA: true})
	// 注册web页面
	app.RegisterView(iris.HTML("./views", ".html"))
	app.Get("/view", func(ctx iris.Context) {
		ctx.View("index.html")
	})
}

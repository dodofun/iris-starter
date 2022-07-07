package router

import (
	"iris-starter/internal/middleware"
	"net/http"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/pprof"
)

func Init(app *iris.Application) {
	// 重复路由报错
	app.SetRegisterRule(iris.RouteError)

	// 引入中间件
	middleware.Init(app)

	// 性能分析
	p := pprof.New()
	app.Any("/debug/pprof", p)
	app.Any("/debug/pprof/{action:path}", p)

	// 注册资源路由
	registerRouter(app)

	// 404
	app.Any("", notFound)
}

func notFound(ctx iris.Context) {
	ctx.StatusCode(http.StatusNotFound)
}

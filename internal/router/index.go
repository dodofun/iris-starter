package router

import (
	"iris-starter/internal/middleware"
	"iris-starter/internal/resources/user"
	"net/http"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/pprof"
	"github.com/kataras/iris/v12/middleware/recover"
)

func Init(app *iris.Application) {
	// 路由不覆盖
	app.SetRegisterRule(iris.RouteSkip)

	// 中间件
	// 从 panics 和 logs恢复
	app.Use(recover.New())
	// 设置请求头
	app.Use((middleware.Header))
	// 用户鉴权
	app.Use((middleware.Auth))

	// 性能分析
	p := pprof.New()
	app.Any("/debug/pprof", p)
	app.Any("/debug/pprof/{action:path}", p)

	// 注册资源路由
	user.RegisterRouter(app)

	// 404
	app.Any("", notFound)
}

func notFound(ctx iris.Context) {
	ctx.StatusCode(http.StatusNotFound)
}

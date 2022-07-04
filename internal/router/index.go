package router

import (
	"iris-starter/internal/middleware"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/pprof"
	"github.com/kataras/iris/v12/middleware/recover"
)

func Init(app *iris.Application) {
	// 中间件
	// 从 panics 和 logs恢复
	app.Use(recover.New())
	// 设置请求头
	app.Use((middleware.Header))
	// 用户鉴权

	// 性能分析
	p := pprof.New()
	app.Any("/debug/pprof", p)
	app.Any("/debug/pprof/{action:path}", p)

	// 注册资源路由

}

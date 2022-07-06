package middleware

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/recover"
)

// 初始化中间件
func Init(app *iris.Application) {
	// 从 panics 和 logs恢复
	app.Use(recover.New())
	// 限速
	// app.Use(Limit)
	// 设置请求头
	app.Use(Header)
	// 用户鉴权
	app.Use(Auth)
	// 压缩
	app.Use(iris.Compression)

}

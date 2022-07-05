package router

import (
	"iris-starter/internal/resources/user"

	"github.com/kataras/iris/v12"
)

// 注册资源路由
func registerRouter(app *iris.Application) {
	// 用户
	user.RegisterRouter(app)
}

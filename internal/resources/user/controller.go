package user

import (
	"github.com/kataras/iris/v12"
)

const (
	basePath = "/v1/user"
	listPath = basePath + "s"
)

var (
	userService Service
)

func init() {
	// 注入依赖
	userService = initService()
}

func RegisterRouter(app *iris.Application) {
	app.Get(basePath, get)
	app.Get(listPath, getList)
	app.Post(basePath, post)
	app.Put(basePath, put)
	app.Delete(basePath, delete)
}

func get(ctx iris.Context) {
	userService.Get()
}

func getList(ctx iris.Context) {
	userService.GetList()
}

func post(ctx iris.Context) {
	userService.Post()
}

func put(ctx iris.Context) {
	userService.Put()
}

func delete(ctx iris.Context) {
	userService.Delete()
}

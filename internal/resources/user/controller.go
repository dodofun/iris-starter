package user

import (
	"github.com/kataras/iris/v12"
)

const (
	basePath = "/v1/user"
	listPath = basePath + "s"
)

var (
	ctl *controller
)

type controller struct {
	service Service
}

func newController(service Service) *controller {
	return &controller{
		service,
	}
}

func Init(app *iris.Application) *controller {
	// 注入依赖
	ctl = initController()
	// 注册路由
	registerRouter(app)
	// 返回 controller
	return ctl
}

func registerRouter(app *iris.Application) {
	app.Get(basePath, get)
	app.Get(listPath, getList)
	app.Post(basePath, post)
	app.Put(basePath, put)
	app.Delete(basePath, delete)
}

func get(ctx iris.Context) {
	ctl.service.get()
}

func getList(ctx iris.Context) {
	ctl.service.getList()
}

func post(ctx iris.Context) {
	ctl.service.post()
}

func put(ctx iris.Context) {
	ctl.service.put()
}

func delete(ctx iris.Context) {
	ctl.service.delete()
}

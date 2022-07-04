package user

import "github.com/kataras/iris/v12"

const (
	basePath = "/v1/user"
	listPath = basePath + "s"
)

func RegisterRouter(app *iris.Application) {
	app.Get(basePath, get)
	app.Get(listPath, getList)
	app.Post(basePath, post)
	app.Put(basePath, put)
	app.Delete(basePath, delete)
}

func get(ctx iris.Context) {

}

func getList(ctx iris.Context) {

}

func post(ctx iris.Context) {

}

func put(ctx iris.Context) {

}

func delete(ctx iris.Context) {

}

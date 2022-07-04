package user

import "github.com/kataras/iris/v12"

const (
	basePath = "/v1/user"
	listPath = basePath + "s"
)

func RegisterRouter(app *iris.Application) {
	party := app.Party(basePath)
	{
		party.Get(basePath, Get)
		party.Get(listPath, GetList)
		party.Post(basePath, Post)
		party.Put(basePath, Put)
		party.Delete(basePath, Delete)
	}
}

func Get(ctx iris.Context) {

}

func GetList(ctx iris.Context) {

}

func Post(ctx iris.Context) {

}

func Put(ctx iris.Context) {

}

func Delete(ctx iris.Context) {

}

package user

import (
	"iris-starter/pkg/logging"
	"iris-starter/pkg/request"

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

// @Summary 获取用户
// @Description 根据用户ID获取用户信息
// @Accept  json
// @Produce  json
// @Param id query int  true  "用户ID"
// @Success 200 {string} string	"ok"
// @Router /v1/user [get]
func get(ctx iris.Context) {
	var query request.QueryId
	err := ctx.ReadQuery(&query)
	logging.Infof("ID: ", query.Id)
	logging.Infof("ERROR: ", err)
	if err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}
	userService.Get()
}

func getList(ctx iris.Context) {
	var query request.QueryPaging
	err := ctx.ReadQuery(&query)
	logging.Infof("Page: ", query.Page)
	logging.Infof("ERROR: ", err)
	if err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}
	userService.GetList()
}

func post(ctx iris.Context) {
	var body PostRequest
	err := ctx.ReadJSON(&body)
	logging.Infof("ID: ", body.Name)
	logging.Infof("ERROR: ", err)
	if err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}
	userService.Post()
}

func put(ctx iris.Context) {
	var body PutRequest
	err := ctx.ReadJSON(&body)
	logging.Infof("ID: ", body.Name)
	logging.Infof("ERROR: ", err)
	if err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}
	userService.Put()
}

func delete(ctx iris.Context) {
	var query request.QueryId
	err := ctx.ReadQuery(&query)
	logging.Infof("ID: ", query.Id)
	logging.Infof("ERROR: ", err)
	if err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}
	userService.Delete()
}

package middleware

import (
	"iris-starter/internal/config"
	"net/http"

	"github.com/kataras/iris/v12"
)

var (
	serverConfig = config.GetConfig().Server
)

type AppHeaders struct {
	// 应用ID
	AppId string `header:"AppId,required"`
	// 平台类型
	PlatformType string `header:"PlatformType,required"`
	// Token
	Token string `header:"Token,required"`
}

// Header 读取/设置请求头.
func Header(ctx iris.Context) {
	// 设置header
	ctx.Header("Access-Control-Allow-Origin", serverConfig.AccessControlAllowOrigin)
	ctx.Header("Access-Control-Allow-Headers", serverConfig.AccessControlAllowHeaders)
	ctx.Header("Access-Control-Allow-Methods", serverConfig.AccessControlAllowMethods)
	if ctx.Request().Method == http.MethodOptions {
		ctx.StatusCode(http.StatusNoContent)
		return
	}

	// 读取header
	var hs AppHeaders
	if err := ctx.ReadHeaders(&hs); err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}
	// TODO 校验/解析header，判断是否有权限访问

	ctx.Next()
}

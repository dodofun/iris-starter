package middleware

import (
	"iris-starter/internal/config"

	"github.com/kataras/iris/v12"
)

var (
	serverConfig = config.GetConfig().Server
)

// Header 设置请求头.
func Header(ctx iris.Context) {
	ctx.Header("Access-Control-Allow-Origin", serverConfig.AccessControlAllowOrigin)
	ctx.Header("Access-Control-Allow-Headers", serverConfig.AccessControlAllowHeaders)
	ctx.Header("Access-Control-Allow-Methods", serverConfig.AccessControlAllowMethods)

	if ctx.Request().Method == "OPTIONS" {
		ctx.StatusCode(204)
		return
	}
	ctx.Next()
}

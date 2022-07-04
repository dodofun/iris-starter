package middleware

import (
	"iris-starter/internal/config"
	"iris-starter/pkg/jwt"
	"iris-starter/pkg/logging"

	"github.com/kataras/iris/v12"
)

// Auth 验证中间件.
func Auth(ctx iris.Context) {

	duration := config.GetConfig().HttpAuthDuration
	// 生成token
	token, err := jwt.CreateToken(map[string]interface{}{"userId": "123456"}, duration)
	if err != nil {
		// TODO 返回异常
		return
	}

	// 获取header中的token
	// token = ctx.GetHeader(jwt.Authorization)
	// if token == "" {
	// 	// TODO 返回异常
	// 	return
	// }

	// 解析token
	value, err := jwt.ParseToken(token)
	if err != nil {
		// TODO 异常处理
		logging.Info("err: " + err.Error())
		return
	}
	logging.Info("value:" + value["userId"].(string))
	ctx.Next()

}

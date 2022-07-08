package docs

import (
	"time"

	"github.com/iris-contrib/swagger"
	"github.com/iris-contrib/swagger/swaggerFiles"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/cache"
)

func RegisterRouter(app *iris.Application) {
	swaggerUI := swagger.Handler(swaggerFiles.Handler,
		swagger.URL("/assets/docs/swagger.yaml"), // The url pointing to API definition.
		swagger.DeepLinking(true),
		swagger.Prefix("/apidoc"),
	)
	// 设置资源缓存，缓存10秒，之后会被清除并重置
	cacheHandler := cache.Handler(1 * time.Second)
	app.Get("/apidoc", cacheHandler, swaggerUI)
	app.Get("/apidoc/{any:path}", swaggerUI)
}

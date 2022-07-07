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
		swagger.URL("http://localhost:8080/swagger/doc.json"), // The url pointing to API definition.
		swagger.DeepLinking(true),
		swagger.Prefix("/swagger"),
	)
	// 设置资源缓存，缓存10秒，之后会被清除并重置
	cacheHandler := cache.Handler(10 * time.Second)
	app.Get("/swagger", cacheHandler, swaggerUI)
	app.Get("/swagger/{any:path}", swaggerUI)
}

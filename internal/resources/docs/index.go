package docs

import (
	"time"

	"github.com/iris-contrib/swagger"
	"github.com/iris-contrib/swagger/swaggerFiles"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/cache"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /v2
func RegisterRouter(app *iris.Application) {
	swaggerUI := swagger.Handler(swaggerFiles.Handler,
		swagger.URL("/assets/doc.yaml"), // The url pointing to API definition.
		swagger.DeepLinking(true),
		swagger.Prefix("/swagger"),
	)
	// 设置资源缓存，缓存10秒，之后会被清除并重置
	cacheHandler := cache.Handler(10 * time.Second)
	app.Get("/swagger", cacheHandler, swaggerUI)
	app.Get("/swagger/{any:path}", swaggerUI)
}

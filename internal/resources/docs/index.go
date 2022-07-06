package docs

import (
	"github.com/iris-contrib/swagger"
	"github.com/iris-contrib/swagger/swaggerFiles"
	"github.com/kataras/iris/v12"
)


func RegisterRouter(app *iris.Application) {
	swaggerUI := swagger.Handler(swaggerFiles.Handler,
		swagger.URL("http://localhost:8080/swagger/doc.json"), // The url pointing to API definition.
		swagger.DeepLinking(true),
		swagger.Prefix("/swagger"),
	)
	app.Get("/swagger", swaggerUI)
	app.Get("/swagger/{any:path}", swaggerUI)
}

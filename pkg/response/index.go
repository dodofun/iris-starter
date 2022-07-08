package response

import "github.com/kataras/iris/v12"

// 错误：参数
func ErrorParam() {

}

// 错误：返回
func Error() {

}

//
func OkObj(ctx iris.Context, data interface{}) {
	ctx.Negotiation().Charset("utf-8")
	ctx.Negotiation().JSON().Binary().Text().EncodingGzip()

	ctx.Negotiate(iris.N{
		Binary: data, // for application/octet-stream,
		JSON:   data, // for application/json
		XML:    data, // for application/xml or text/xml
		// Other: for anything else,
	})
}

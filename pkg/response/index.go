package response

import (
	"github.com/kataras/iris/v12"
	"google.golang.org/protobuf/proto"
)

// 错误：参数
func ErrorParam() {

}

// 错误：返回
func ErrorBadRequest(ctx iris.Context, err error) {
	ctx.StopWithError(iris.StatusBadRequest, err)
}

//
func OkObj(ctx iris.Context, data proto.Message) {
	ctx.Negotiation().Charset("utf-8")
	ctx.StatusCode(iris.StatusOK)

	// for application/json
	options := iris.JSON{
		Proto: iris.ProtoMarshalOptions{
			AllowPartial: true,
			Multiline:    true,
			Indent:       "    ",
		},
	}

	// protobuf: application/x-protobuf
	// json: application/json
	ctx.Negotiation().JSON(data, options).Protobuf(data).EncodingGzip()
	ctx.Negotiation().Accept.JSON() //  默认返回json

	err := ctx.Err()
	if err != nil {
		ErrorBadRequest(ctx, err)
		return
	}
}

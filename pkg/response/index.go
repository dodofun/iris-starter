package response

import (
	"github.com/kataras/iris/v12"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

// 错误：参数错误
func ErrorParam(ctx iris.Context, err error) {
	ctx.StopWithError(iris.StatusBadRequest, err)
}

// 错误：发生异常
func ErrorBadRequest(ctx iris.Context, err error) {
	ctx.StopWithError(iris.StatusBadRequest, err)
}

// 正常：单条数据
func OkObj(ctx iris.Context, data proto.Message) {
	OkObjWithMsg(ctx, data, HeaderMsgSuccess)
}

func OkObjWithMsg(ctx iris.Context, data proto.Message, msg string) {
	OkObjWithMsgAndStatusCode(ctx, data, msg, iris.StatusOK)
}

func OkObjWithMsgAndStatusCode(ctx iris.Context, data proto.Message, msg string, statusCode int) {

	header := &Header{
		Code:    HeaderCodeSuccess,
		Message: msg,
	}

	anypb.New(data)

	any, err := anypb.New(data)
	if err != nil {
		ErrorBadRequest(ctx, err)
		return
	}

	obj := &Obj{
		Header: header,
		Data:   any,
	}

	// 默认返回json
	ctx.Negotiation().Accept.JSON()
	// 根据 header 的 accept 值不同，返回不同格式数据，没有设置accept时，返回json格式数据
	// protobuf: application/x-protobuf
	// json: application/json
	ctx.Negotiation().Charset("utf-8").JSON().Protobuf().EncodingGzip()

	ctx.Negotiate(iris.N{
		JSON:     obj, // for application/json
		Protobuf: obj, // for application/x-protobuf
	})

	// response 状态码
	if statusCode > 0 {
		ctx.StatusCode(statusCode)
	}

	err = ctx.Err()
	if err != nil {
		ErrorBadRequest(ctx, err)
		return
	}
}

// 正常：列表数据
func OkList(ctx iris.Context, data []proto.Message, paging *Paging) {
	OkListWithMsg(ctx, data, paging, HeaderMsgSuccess)
}
func OkListWithMsg(ctx iris.Context, data []proto.Message, paging *Paging, msg string) {
	OkListWithMsgAndStatusCode(ctx, data, paging, msg, iris.StatusOK)
}

func OkListWithMsgAndStatusCode(ctx iris.Context, data []proto.Message, paging *Paging, msg string, statusCode int) {

	header := &Header{
		Code:    0,
		Message: msg,
	}

	anyDatas := []*anypb.Any{}
	for _, v := range data {
		any, _ := anypb.New(v)
		anyDatas = append(anyDatas, any)
	}

	obj := &List{
		Header: header,
		Paging: paging,
		Data:   anyDatas,
	}

	// 默认返回json
	ctx.Negotiation().Accept.JSON()
	// 根据 header 的 accept 值不同，返回不同格式数据，没有设置accept时，返回json格式数据
	// protobuf: application/x-protobuf
	// json: application/json
	ctx.Negotiation().Charset("utf-8").JSON().Protobuf().EncodingGzip()

	ctx.Negotiate(iris.N{
		JSON:     obj, // for application/json
		Protobuf: obj, // for application/x-protobuf
	})

	// response 状态码
	if statusCode > 0 {
		ctx.StatusCode(statusCode)
	}

	err := ctx.Err()
	if err != nil {
		ErrorBadRequest(ctx, err)
		return
	}
}

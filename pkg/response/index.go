package response

import (
	"github.com/kataras/iris/v12"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

// 错误：参数
func ErrorParam() {

}

// 错误：返回
func ErrorBadRequest(ctx iris.Context, err error) {
	ctx.StopWithError(iris.StatusBadRequest, err)
}

// 正常：单条数据
func OkObj(ctx iris.Context, data *proto.Message, msg string) {
	OkObjWithStatusCode(ctx, data, msg, iris.StatusOK)
}

func OkObjWithStatusCode(ctx iris.Context, data *proto.Message, msg string, statusCode int) {

	header := &Header{
		Code:    0,
		Message: msg,
	}

	anypb.New(*data)

	any, err := anypb.New(*data)
	if err != nil {
		ErrorBadRequest(ctx, err)
		return
	}

	obj := &Obj{
		Header: header,
		Data:   any,
	}

	options := iris.JSON{
		Proto: iris.ProtoMarshalOptions{
			AllowPartial: true,
			Multiline:    true,
			Indent:       "    ",
		},
	}

	// 根据 header 的 accept 值不同，返回不同格式数据，没有设置accept时，返回json格式数据
	// protobuf: application/x-protobuf
	// json: application/json
	ctx.Negotiation().Charset("utf-8").JSON(obj, options).Protobuf(obj).EncodingGzip()
	//  默认返回json
	ctx.Negotiation().Accept.JSON()
	// response 状态码
	if statusCode > 0 {
		ctx.StatusCode(statusCode)
	} else {
		ctx.StatusCode(iris.StatusOK)
	}

	err = ctx.Err()
	if err != nil {
		ErrorBadRequest(ctx, err)
		return
	}
}

// 正常：列表数据
func OkList(ctx iris.Context, data *[]proto.Message, paging *Paging, msg string) {
	OkListWithStatusCode(ctx, data, paging, msg, iris.StatusOK)
}
func OkListWithStatusCode(ctx iris.Context, data *[]proto.Message, paging *Paging, msg string, statusCode int) {

	header := &Header{
		Code:    0,
		Message: msg,
	}

	anyDatas := []*anypb.Any{}
	for _, v := range *data {
		any, _ := anypb.New(v)
		anyDatas = append(anyDatas, any)
	}

	obj := &List{
		Header: header,
		Paging: paging,
		Data:   anyDatas,
	}

	options := iris.JSON{
		Proto: iris.ProtoMarshalOptions{
			AllowPartial: true,
			Multiline:    true,
			Indent:       "    ",
		},
	}

	// 根据 header 的 accept 值不同，返回不同格式数据，没有设置accept时，返回json格式数据
	// protobuf: application/x-protobuf
	// json: application/json
	ctx.Negotiation().Charset("utf-8").JSON(obj, options).Protobuf(obj).EncodingGzip()
	//  默认返回json
	ctx.Negotiation().Accept.JSON()
	// response 状态码
	if statusCode > 0 {
		ctx.StatusCode(statusCode)
	} else {
		ctx.StatusCode(iris.StatusOK)
	}

	err := ctx.Err()
	if err != nil {
		ErrorBadRequest(ctx, err)
		return
	}
}

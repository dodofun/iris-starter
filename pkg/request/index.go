package request

type QueryId struct {
	Id uint64 `url:"id" validate:"required"`
}

type QueryPaging struct {
	Page     uint64 `url:"page"`
	PageSize uint64 `url:"pageSize"`
}

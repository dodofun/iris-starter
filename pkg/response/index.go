package response

type Obj[T any] struct {
	Header Header `json:"header"`
	Data   T      `json:"data"`
}

type List[T any] struct {
	Header Header `json:"header"`
	Paging Paging `json:"paging"`
	Data   []T    `json:"data"`
}

type Header struct {
	Code    uint64 `json:"code"`
	Message string `json:"message"`
}

type Paging struct {
	Page     uint64 `json:"page"`
	PageSize uint64 `json:"pageSize"`
	Total    uint64 `json:"total"`
}

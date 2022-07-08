package user

type PostRequest struct {
	Name string `json:"name" validate:"required"`
	Mobile uint64   `json:"age" validate:"gte=0,lte=130"`
}

type PutRequest struct {
	Id uint64
	Name string `json:"name" validate:"required"`
	Mobile uint64   `json:"age" validate:"gte=0,lte=130"`
}

type User struct {
	Id uint64
	Name string
	Mobile uint64
	CreateAt uint64
	UpdateAt uint64
	Enable bool
}
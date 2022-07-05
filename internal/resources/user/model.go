package user

type PostRequest struct {
	Name string `json:"name" validate:"required"`
	Age  uint   `json:"age" validate:"gte=0,lte=130"`
}

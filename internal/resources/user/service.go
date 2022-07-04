package user

type Service interface {
	Get()
	GetList()
	Post()
	Put()
	Delete()
}

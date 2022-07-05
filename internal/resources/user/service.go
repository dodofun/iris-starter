package user

type Service interface {
	Get()
	GetList()
	Post()
	Put()
	Delete()
}

type service struct {
	dao   Dao
	cache Cache
}

func newService(dao Dao, cache Cache) Service {
	return &service{dao, cache}
}

func (s *service) Get() {

}

func (s *service) GetList() {

}

func (s *service) Post() {

}

func (s *service) Put() {

}

func (s *service) Delete() {

}

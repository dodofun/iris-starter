package user

type Dao interface {
	Get()
	GetList()
	Post()
	Put()
	Delete()
}
type dao struct {
}

func newDao() Dao {
	return &dao{}
}

func (s *dao) Get() {

}

func (s *dao) GetList() {

}

func (s *dao) Post() {

}

func (s *dao) Put() {

}

func (s *dao) Delete() {

}
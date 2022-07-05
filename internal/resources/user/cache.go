package user

type Cache interface {
	Get()
	GetList()
	Post()
	Put()
	Delete()
}
type cache struct {
}

func newCache() Cache {
	return &cache{}
}

func (s *cache) Get() {

}

func (s *cache) GetList() {

}

func (s *cache) Post() {

}

func (s *cache) Put() {

}

func (s *cache) Delete() {

}

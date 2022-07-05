package user

type Service interface {
	get()
	getList()
	post()
	put()
	delete()
}

type service struct {
	dao   Dao
	cache Cache
}

func newService(dao Dao, cache Cache) Service {
	return &service{dao, cache}
}

func (s *service) get() {

}

func (s *service) getList() {

}

func (s *service) post() {

}

func (s *service) put() {

}

func (s *service) delete() {

}

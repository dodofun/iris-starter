package user

type Dao interface {
	get()
	getList()
	post()
	put()
	delete()
}
type dao struct {
}

func newDao() Dao {
	return &dao{}
}

func (s *dao) get() {

}

func (s *dao) getList() {

}

func (s *dao) post() {

}

func (s *dao) put() {

}

func (s *dao) delete() {

}

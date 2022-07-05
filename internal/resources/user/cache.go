package user

type Cache interface {
	get()
	getList()
	post()
	put()
	delete()
}
type cache struct {
}

func newCache() Cache {
	return &cache{}
}

func (s *cache) get() {

}

func (s *cache) getList() {

}

func (s *cache) post() {

}

func (s *cache) put() {

}

func (s *cache) delete() {

}

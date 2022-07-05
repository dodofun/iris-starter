package user

import "github.com/google/wire"

type Service struct {
	dao dao
}

func newService(dao dao) Service {
	return Service{dao: dao}
}

func InitService() (Service, error) {
	wire.Build(newService, newDao)
	return Service{}, nil
}

func (s *Service) get() {

}

func (s *Service) getList() {

}

func (s *Service) post() {

}

func (s *Service) put() {

}

func (s *Service) delete() {

}

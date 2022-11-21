package server

import (
	hdl "github.com/vietpham1023/golang-uit-hackathon/handler"
	"github.com/vietpham1023/golang-uit-hackathon/internal/repos"
	"github.com/vietpham1023/golang-uit-hackathon/internal/services/book"
)

func (s *Server) initServices(repo repos.IRepo) *ServiceList {
	book := book.NewBook(s.cfg, repo)
	return &ServiceList{
		book: book,
	}
}

func (s *Server) initRouters(serviceList *ServiceList) {

	handler := hdl.NewHandler(
		s.cfg,
		s.router,
		serviceList.book,
	)

	handler.ConfigureAPIRoute(s.router)

}

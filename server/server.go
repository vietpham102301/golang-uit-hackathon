package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/vietpham1023/golang-uit-hackathon/config"
	"github.com/vietpham1023/golang-uit-hackathon/infra"
	"github.com/vietpham1023/golang-uit-hackathon/internal/repos"
	"github.com/vietpham1023/golang-uit-hackathon/internal/services/book"
	"go.uber.org/zap"
	"net/http"
	"os"
)

type Server struct {
	httpServer   *http.Server
	router       *gin.Engine
	metricRouter *gin.Engine
	cfg          *config.AppConfig
}

type ServiceList struct {
	book book.IBook
}

func NewServer(cfg *config.AppConfig) (*Server, error) {
	router := gin.New()
	router.Use(gin.Recovery())

	s := &Server{
		router: router,
		cfg:    cfg,
	}

	return s, nil
}

func (s *Server) Init() {
	db, err := infra.InitPostgres(s.cfg.DB)
	if err != nil {
		zap.S().Errorf("Init db error: %v", err)
		panic(err)
	}

	repo := repos.NewSQLRepo(db)

	serviceList := s.initServices(repo)
	s.initRouters(serviceList)
}

func (s *Server) Listen() error {
	address := fmt.Sprintf(":%s", os.Getenv("PORT"))
	fmt.Sprintf(":%s", os.Getenv("PORT"))
	s.httpServer = &http.Server{
		Handler: s.router,
		Addr:    address,
	}

	return s.httpServer.ListenAndServe()
}
package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/vietpham1023/golang-uit-hackathon/config"
	"github.com/vietpham1023/golang-uit-hackathon/internal/services/book"
	"net/http"
)

type Handler struct {
	cfg    *config.AppConfig
	router *gin.Engine
	book   book.IBook
}

func NewHandler(
	cfg *config.AppConfig,
	router *gin.Engine,
	book book.IBook,
) *Handler {
	return &Handler{
		cfg:    cfg,
		router: router,
		book:   book,
	}
}

func (h *Handler) ConfigureAPIRoute(router *gin.Engine) {
	routers := router.Group("v1")
	routers.POST("/book", h.createBook())
	routers.GET("/smt", func(c *gin.Context) {
		c.JSON(http.StatusOK, "hello world!")
	})
}

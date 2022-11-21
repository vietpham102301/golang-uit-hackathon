package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/vietpham1023/golang-uit-hackathon/handler/models"
	models2 "github.com/vietpham1023/golang-uit-hackathon/internal/models"
	"net/http"
)

func (h *Handler) createBook() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("start creating book")
		input := models.CreateBookRequest{}
		err := ctx.Bind(&input)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		record := &models2.Book{}
		err = copier.Copy(record, input)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		record, err = h.book.CreateBook(record)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		ctx.JSON(http.StatusOK, record)
	}
}

package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/vietpham1023/golang-uit-hackathon/handler/models"
	models2 "github.com/vietpham1023/golang-uit-hackathon/internal/models"
	"net/http"
	"strconv"
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

func (h *Handler) GetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("get by book id")
		id, ok := ctx.GetQuery("id")
		if !ok {
			ctx.JSON(http.StatusBadRequest, "query param invalid!")
			return
		}
		iID, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		record, err := h.book.GetByID(iID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		ctx.JSON(http.StatusOK, record)
	}
}

func (h *Handler) listBookByFilter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("get book by filter...")
		q, data := ctx.Request.URL.Query(), map[string]string{}
		for k, v := range q {
			if len(v) > 0 {
				data[k] = v[0]
			}
		}

		records, err := h.book.ListByFilter(data)
		if err != nil {
			fmt.Printf("list book fail with err: %v\n", err)
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, records)
	}
}

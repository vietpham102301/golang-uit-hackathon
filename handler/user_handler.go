package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/vietpham1023/golang-uit-hackathon/handler/models"
	models2 "github.com/vietpham1023/golang-uit-hackathon/internal/models"
	"net/http"
)

func (h *Handler) createUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("start creating user")
		input := models.CreateUserRequest{}
		err := ctx.Bind(&input)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		record := &models2.User{}
		err = copier.Copy(record, input)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		record, err = h.user.CreateUser(record)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		res := &models.CreateUserResponse{}
		ctx.JSON(http.StatusOK, res.ToUserResponse(record))
	}
}

func (h *Handler) logInUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("start login user")
		input := models.CreateUserRequest{}
		err := ctx.Bind(&input)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		record := &models2.User{}
		err = copier.Copy(record, input)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		record, err = h.user.Login(record)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		res := &models.CreateUserResponse{}
		ctx.JSON(http.StatusOK, res.ToUserResponse(record))
	}
}

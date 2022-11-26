package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) APIAuthentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("X-Access-Token")

		user2, err := h.user.GetUserByToken(token)
		if err != nil {
			fmt.Printf("get user by token fail with err: %v\n", err)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthorized")
		}
		ctx.Set("user_id", user2.ID)
		ctx.Next()
	}
}

package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) getItemsByProviderID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("get GetItemsByProviderID...")
		id, ok := ctx.GetQuery("provider_id")
		if !ok {
			ctx.JSON(http.StatusBadRequest, "query param invalid!")
			return
		}
		iID, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		res, err := h.item.GetItemsByProviderID(iID)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			return
		}

		if err != nil {
			fmt.Printf("list merchant-campaign fail with err: %v\n", err)
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, res)
	}
}

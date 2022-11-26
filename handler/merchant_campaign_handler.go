package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) listMerchantCampaignByFilter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("get merchant CAMPAIGN by filter...")
		q, data := ctx.Request.URL.Query(), map[string]string{}
		for k, v := range q {
			if len(v) > 0 {
				data[k] = v[0]
			}
		}

		res, err := h.merchantCampaign.ListByFilter(data)
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

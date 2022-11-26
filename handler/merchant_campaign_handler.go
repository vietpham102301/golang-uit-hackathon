package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	handlerModels "github.com/vietpham1023/golang-uit-hackathon/handler/models"
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
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		err = h.merchantCampaign.Create(ctx, merchantCampaign)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		ctx.JSON(http.StatusOK, res)
	}
}

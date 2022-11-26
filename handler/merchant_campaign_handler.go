package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) listMerchantCampaignByFilter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("get book by filter...")
		q, data := ctx.Request.URL.Query(), map[string]string{}
		for k, v := range q {
			if len(v) > 0 {
				data[k] = v[0]
			}
		}

		records, err := h.merchantCampaign.ListByFilter(data)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			return
		}
		var campaignIDs []int64

		for _, record := range records {
			campaignIDs = append(campaignIDs, record.CampaignID)
		}
		campaigns, err := h.campaign.ListCampaignByIDs(campaignIDs)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			return
		}

		if err != nil {
			fmt.Printf("list merchant-campaign fail with err: %v\n", err)
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, records)
	}
}

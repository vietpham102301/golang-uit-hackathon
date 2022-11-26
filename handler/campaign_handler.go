package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/vietpham1023/golang-uit-hackathon/handler/models"
	models2 "github.com/vietpham1023/golang-uit-hackathon/internal/models"
	"net/http"
)

func (h *Handler) createCampaign() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("start creating campaign!")
		input := models.CampaignRequest{}
		err := ctx.Bind(&input)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		rule, err := h.rule.CreateRule(&models2.Rule{
			ItemId:        input.Rule.ItemID,
			RequiredObNum: input.Rule.RequiredObNum,
		})

		if err != nil {
			fmt.Println("err: %v", err)
			return
		}

		record := &models2.Campaign{}
		err = copier.Copy(record, input)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		record.RuleID = rule.ID
		userID, ok := ctx.Get("user_id")
		if ok {
			record.ProviderID = userID.(int64)
		}
		record, err = h.campaign.CreateCampaign(record)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		ctx.JSON(http.StatusOK, record)
	}
}

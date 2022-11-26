package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/vietpham1023/golang-uit-hackathon/handler/models"
	"net/http"
)

func (h *Handler) createMerchant() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var merchant *models.Merchant
		err := ctx.ShouldBind(&merchant)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		err = h.merchant.CreateMerchant(ctx, merchant)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		ctx.JSON(http.StatusOK, nil)
		return
	}
}

func (h *Handler) getMerchantByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var condition models.GetMerchant

		err := ctx.ShouldBindUri(&condition)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		id := condition.ID

		merchant, err := h.merchant.GetByID(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, merchant)
		return
	}
}

func (h *Handler) ListMerchantByCondition() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var merchant models.ListMerchant
		err := ctx.ShouldBindQuery(&merchant)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		mapMerchant := map[string]interface{}{
			"id":   merchant.ID,
			"name": merchant.Name,
		}
		merchants, err := h.merchant.ListByFilter(ctx, mapMerchant)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		ctx.JSON(http.StatusOK, merchants)
		return
	}
}

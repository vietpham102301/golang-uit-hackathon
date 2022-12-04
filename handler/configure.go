package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/vietpham1023/golang-uit-hackathon/config"
	"github.com/vietpham1023/golang-uit-hackathon/internal/services/book"
	"github.com/vietpham1023/golang-uit-hackathon/internal/services/campaign"
	item2 "github.com/vietpham1023/golang-uit-hackathon/internal/services/item"
	"github.com/vietpham1023/golang-uit-hackathon/internal/services/merchant"
	"github.com/vietpham1023/golang-uit-hackathon/internal/services/merchant_campaign"
	"github.com/vietpham1023/golang-uit-hackathon/internal/services/provider"
	rule2 "github.com/vietpham1023/golang-uit-hackathon/internal/services/rule"
	user2 "github.com/vietpham1023/golang-uit-hackathon/internal/services/user"
	"log"
	"net/http"
	"path/filepath"
)

type Handler struct {
	cfg              *config.AppConfig
	router           *gin.Engine
	book             book.IBook
	user             user2.IUser
	merchantCampaign merchant_campaign.IMerchantCampaign
	campaign         campaign.ICampaign
	rule             rule2.IRule
	item             item2.IItem
	merchant         merchant.IMerchant
	provider         provider.IProvider
}

func NewHandler(
	cfg *config.AppConfig,
	router *gin.Engine,
	book book.IBook,
	user user2.IUser,
	merchantCampaign merchant_campaign.IMerchantCampaign,
	campaign campaign.ICampaign,
	rule rule2.IRule,
	item item2.IItem,
	merchant merchant.IMerchant,
	provider provider.IProvider,
) *Handler {
	return &Handler{
		cfg:              cfg,
		router:           router,
		book:             book,
		user:             user,
		merchantCampaign: merchantCampaign,
		campaign:         campaign,
		rule:             rule,
		item:             item,
		merchant:         merchant,
		provider:         provider,
	}
}

func (h *Handler) ConfigureAPIRoute(router *gin.Engine) {

	// v1 for user
	routers := router.Group("v1")
	//routers.GET("/book", h.GetByID())
	//routers.POST("/book", h.createBook())
	//routers.GET("/book", h.GetByID())

	routers.GET("/file", func(ctx *gin.Context) {
		ctx.File("savedfile/65300bcbc0c4029a5bd5.jpg")
	})
	routers.POST("/upload", func(ctx *gin.Context) {
		// single file
		file, _ := ctx.FormFile("file")
		log.Println(file.Filename)

		// Retrieve file information
		extension := filepath.Ext(file.Filename)
		// Generate random file name for the new uploaded file so it doesn't override the old file with same name
		newFileName := uuid.New().String() + extension

		// Upload the file to specific dst.
		err := ctx.SaveUploadedFile(file, "savedfile/"+newFileName+extension)
		if err != nil {
			fmt.Printf("error when save file!!: %v\n", err)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Unable to save the file smt",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"data":   file,
		})
	})

	routers.POST("/signup", h.createUser())
	routers.POST("/login", h.logInUser())
	routers.Use(h.APIAuthentication())
	//routers.GET("/book/list", h.listBookByFilter())

	//Merchant
	routers.GET("/merchant-campaign/list", h.listMerchantCampaignByFilter())
	routers.GET("/item/list", h.getItemsByProviderID())
	routers.POST("/campaign", h.createCampaign())
	routers.GET("/merchant/:id", h.getMerchantByID())
	//routers.POST("merchant-campaign", h.cre)
}

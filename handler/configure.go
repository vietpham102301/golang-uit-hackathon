package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/vietpham1023/golang-uit-hackathon/config"
	"github.com/vietpham1023/golang-uit-hackathon/internal/services/book"
	"github.com/vietpham1023/golang-uit-hackathon/internal/services/campaign"
	merchantService "github.com/vietpham1023/golang-uit-hackathon/internal/services/merchant"
	"github.com/vietpham1023/golang-uit-hackathon/internal/services/merchant_campaign"
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
	merchant         merchantService.IMerchant
}

func NewHandler(
	cfg *config.AppConfig,
	router *gin.Engine,
	book book.IBook,
	user user2.IUser,
	merchantCampaign merchant_campaign.IMerchantCampaign,
	campaign campaign.ICampaign,
	merchant merchantService.IMerchant,
) *Handler {
	return &Handler{
		cfg:              cfg,
		router:           router,
		book:             book,
		user:             user,
		merchantCampaign: merchantCampaign,
		campaign:         campaign,
		merchant:         merchant,
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
				"message": "Unable to save the file",
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
	//routers.Use(h.APIAuthentication())
	//routers.GET("/book/list", h.listBookByFilter())
	routers.GET("merchant-campaign/list", h.listMerchantCampaignByFilter())

}

func (h *Handler) ConfigureAPIMerchantRoute(r *gin.Engine) {

	v1 := r.Group("v1")
	merchant := v1.Group("/merchant")
	{
		merchant.POST("", h.createMerchant())
		merchant.GET("/:id", h.getMerchantByID())
		merchant.GET("", h.ListMerchantByCondition())
	}

	merchantCampaign := v1.Group("/merchant-campaign")
	{
		merchantCampaign.POST("", h.createMerchantCampaign())
	}
}

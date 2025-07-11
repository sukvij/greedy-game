package campaign

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CampaignController struct {
	Db       *gorm.DB
	Campaign *Campaign
}

func CampaignServiceController(app *gin.RouterGroup, db *gorm.DB) {
	campaignController := &CampaignController{Db: db}
	router := app.Group("/campaign")
	router.GET("", campaignController.getAllCampaign)
	router.POST("", campaignController.createCampaign)
	router.PATCH("", campaignController.updateCampaign)
}

func (campaignController *CampaignController) getAllCampaign(ctx *gin.Context) {
	campaignController.Campaign = &Campaign{}
	campaignService := NewCampaignService(campaignController.Db, campaignController.Campaign)
	res, err := campaignService.GetAllCampaign()
	if err != nil {
		ctx.JSON(500, err)
		return
	}
	ctx.JSON(200, res)
}

func (campaignController *CampaignController) createCampaign(ctx *gin.Context) {
	campaignController.Campaign = &Campaign{}
	ctx.BindJSON(campaignController.Campaign)
	campaignService := NewCampaignService(campaignController.Db, campaignController.Campaign)
	res, err := campaignService.CreateCampaign()
	fmt.Println("create cmpaign err bro ", err)
	if err != nil {
		ctx.JSON(500, err)
		return
	}
	ctx.JSON(200, res)
}

func (campaignController *CampaignController) updateCampaign(ctx *gin.Context) {
	campaignController.Campaign = &Campaign{}
	ctx.BindJSON(campaignController.Campaign)
	campaignService := NewCampaignService(campaignController.Db, campaignController.Campaign)
	fmt.Println("good bro")
	res, err := campaignService.UpdateCampaign()
	fmt.Println("update cmpaign err bro ", err)
	if err != nil {
		ctx.JSON(500, err)
		return
	}
	ctx.JSON(200, res)
}

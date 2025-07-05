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

func CampaignServiceController(app *gin.Engine, db *gorm.DB) {
	campaignController := &CampaignController{Db: db, Campaign: &Campaign{}}
	router := app.Group("/campaign")
	router.GET("", campaignController.getAllCampaign)
	router.POST("", campaignController.createCampaign)
}

func (campaignController *CampaignController) getAllCampaign(ctx *gin.Context) {
	campaignService := NewCampaignService(campaignController.Db, campaignController.Campaign)
	res, err := campaignService.GetAllCampaign()
	if err != nil {
		ctx.JSON(500, err)
		return
	}
	ctx.JSON(200, res)
}

func (campaignController *CampaignController) createCampaign(ctx *gin.Context) {
	ctx.BindJSON(campaignController.Campaign)
	fmt.Println(*(campaignController.Campaign), campaignController.Campaign)
	campaignService := NewCampaignService(campaignController.Db, campaignController.Campaign)
	res, err := campaignService.CreateCampaign()
	fmt.Println("create cmpaign err bro ", err)
	ctx.JSON(200, res)
}

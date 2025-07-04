package servicediscovery

import (
	"github.com/gin-gonic/gin"
	"github.com/sukvij/greedy-game/campaign"
	"gorm.io/gorm"
)

func RouteService(app *gin.Engine, db *gorm.DB) {
	campaign.CampaignServiceController(app, db)
}

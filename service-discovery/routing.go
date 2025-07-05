package servicediscovery

import (
	"github.com/gin-gonic/gin"
	"github.com/sukvij/greedy-game/campaign"
	targetingrule "github.com/sukvij/greedy-game/targeting-rule"
	"gorm.io/gorm"
)

func RouteService(app *gin.Engine, db *gorm.DB) {
	campaign.CampaignServiceController(app, db)
	targetingrule.TargetingRuleServiceController(app, db)
}

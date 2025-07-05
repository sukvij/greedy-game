package servicediscovery

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/sukvij/greedy-game/campaign"
	"github.com/sukvij/greedy-game/delivery"
	targetingrule "github.com/sukvij/greedy-game/targeting-rule"
	"gorm.io/gorm"
)

func RouteService(app *gin.Engine, db *gorm.DB, redisClient *redis.Client) {
	campaign.CampaignServiceController(app, db)
	targetingrule.TargetingRuleServiceController(app, db)
	delivery.DeliveryServiceController(app, db, redisClient)
}

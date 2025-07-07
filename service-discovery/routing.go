package servicediscovery

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/sukvij/greedy-game/campaign"
	"github.com/sukvij/greedy-game/delivery"
	"github.com/sukvij/greedy-game/gredfers/logs"
	targetingrule "github.com/sukvij/greedy-game/targeting-rule"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"gorm.io/gorm"
)

func RouteService(app *gin.Engine, db *gorm.DB, redisClient *redis.Client, logs *logs.AgreeGateLoager, tracker *sdktrace.TracerProvider) {
	campaign.CampaignServiceController(app, db)
	targetingrule.TargetingRuleServiceController(app, db)
	delivery.DeliveryServiceController(app, db, redisClient, logs, tracker)
}

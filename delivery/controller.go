package delivery

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	redisservice "github.com/sukvij/greedy-game/redis-service"
	"gorm.io/gorm"
)

type DeliveryController struct {
	Db          *gorm.DB
	Request     *Request
	RedisClient *redis.Client
}

func DeliveryServiceController(app *gin.Engine, db *gorm.DB, redisClient *redis.Client) {
	controller := &DeliveryController{Db: db, RedisClient: redisClient}
	router := app.Group("/delivery")
	router.GET("", controller.getDelivery)
}

func (controller *DeliveryController) getDelivery(ctx *gin.Context) {
	controller.Request = &Request{}
	ctx.Bind(controller.Request)

	key := controller.Request.AppId + "-" + controller.Request.Country + "-" + controller.Request.OperatingStstem
	// fetch result from redis if exist
	redisResult, err1 := redisservice.GetValue(controller.RedisClient, key)
	if err1 == redis.Nil {
		ctx.JSON(200, redisResult)
		return
	}
	fmt.Println("bhai database me jaa rha h abhi to..")
	deliveryService := NewDeliveryService(controller.Db, controller.Request)
	res, _ := deliveryService.GetDelivery()
	redisservice.SetValue(controller.RedisClient, key, res)
	ctx.JSON(200, res)
}

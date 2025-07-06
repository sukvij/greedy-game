package delivery

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	redisservice "github.com/sukvij/greedy-game/redis-service"
	"github.com/sukvij/greedy-game/response"
	"gorm.io/gorm"
)

type DeliveryController struct {
	Db            *gorm.DB
	Request       *Request
	RedisClient   *redis.Client
	APICalledTime time.Time
}

func DeliveryServiceController(app *gin.Engine, db *gorm.DB, redisClient *redis.Client) {
	controller := &DeliveryController{Db: db, RedisClient: redisClient}
	router := app.Group("/delivery")
	router.GET("", controller.getDelivery)
}

func (controller *DeliveryController) getDelivery(ctx *gin.Context) {
	controller.APICalledTime = time.Now()
	controller.Request = &Request{}

	bindErr := ctx.Bind(controller.Request)
	if bindErr != nil {
		response.JSONResponse(ctx, bindErr, nil, time.Since(controller.APICalledTime).Milliseconds())
		return
	}

	key := controller.Request.AppId + "-" + controller.Request.Country + "-" + controller.Request.OperatingStstem
	// fetch result from redis if exist
	redisResult, err1 := redisservice.GetValue(controller.RedisClient, key)
	if err1 == redis.Nil && redisResult != nil {
		response.JSONResponse(ctx, nil, redisResult, time.Since(controller.APICalledTime).Milliseconds())
		return
	}
	fmt.Println("bhai database me jaa rha h abhi to..")
	deliveryService := NewDeliveryService(controller.Db, controller.Request)
	res, err := deliveryService.GetDelivery()
	if err != nil {
		response.JSONResponse(ctx, err, nil, time.Since(controller.APICalledTime).Milliseconds())
		return
	}

	redisservice.SetValue(controller.RedisClient, key, res)
	response.JSONResponse(ctx, nil, res, time.Since(controller.APICalledTime).Milliseconds())
}

package delivery

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/sukvij/greedy-game/gredfers/logs"
	redisservice "github.com/sukvij/greedy-game/gredfers/redis-service"
	"github.com/sukvij/greedy-game/gredfers/response"
	"go.opentelemetry.io/otel"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"gorm.io/gorm"
)

type DeliveryController struct {
	Db            *gorm.DB
	Request       *Request
	RedisClient   *redis.Client
	APICalledTime time.Time
	Loager        *logs.AgreeGateLoager
	Tracker       *sdktrace.TracerProvider
}

func DeliveryServiceController(app *gin.Engine, db *gorm.DB, redisClient *redis.Client, logs *logs.AgreeGateLoager, tracker *sdktrace.TracerProvider) {
	controller := &DeliveryController{Db: db, RedisClient: redisClient, Loager: logs, Tracker: tracker}
	router := app.Group("/delivery")
	router.GET("", controller.getDelivery)
}

func (controller *DeliveryController) getDelivery(ctx *gin.Context) {
	_, span := otel.Tracer("repository").Start(ctx, "NewRepository")
	defer span.End()
	controller.APICalledTime = time.Now()
	controller.Request = &Request{}
	bindErr := ctx.Bind(controller.Request)
	if bindErr != nil {
		controller.Loager.Error(bindErr)
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
		controller.Loager.Error(err)
		response.JSONResponse(ctx, err, nil, time.Since(controller.APICalledTime).Milliseconds())
		return
	}

	redisservice.SetValue(controller.RedisClient, key, res)
	response.JSONResponse(ctx, nil, res, time.Since(controller.APICalledTime).Milliseconds())
}

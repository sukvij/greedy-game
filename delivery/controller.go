package delivery

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DeliveryController struct {
	Db      *gorm.DB
	Request *Request
}

func DeliveryServiceController(app *gin.Engine, db *gorm.DB) {
	controller := &DeliveryController{Db: db}
	router := app.Group("/delivery")
	router.GET("", controller.getDelivery)
}

func (controller *DeliveryController) getDelivery(ctx *gin.Context) {
	controller.Request = &Request{}
	ctx.Bind(controller.Request)
	fmt.Println("before", controller.Request)
	fmt.Println("after", controller.Request)

	deliveryService := NewDeliveryService(controller.Db, controller.Request)
	res, _ := deliveryService.GetDelivery()
	ctx.JSON(200, res)
}

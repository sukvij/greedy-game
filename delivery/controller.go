package delivery

import (
	"fmt"
	"strings"

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
	controller.Request.AppId = strings.Split(controller.Request.AppId[0], ".")
	controller.Request.Country = strings.Split(controller.Request.Country[0], ".")
	controller.Request.OperatingStstem = strings.Split(controller.Request.Country[0], ".")
	fmt.Println("after", controller.Request)

	deliveryService := NewDeliveryService(controller.Db, controller.Request)
	res, _ := deliveryService.GetDelivery()
	ctx.JSON(200, res)
}

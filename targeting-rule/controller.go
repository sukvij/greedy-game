package targetingrule

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TargetingRuleController struct {
	Db            *gorm.DB
	TargetingRule *TargetingRule
}

func TargetingRuleServiceController(app *gin.RouterGroup, db *gorm.DB) {
	targetingRuleController := &TargetingRuleController{Db: db}
	router := app.Group("/targeting-rule")
	router.GET("", targetingRuleController.getAllTargetingRule)
	router.PATCH("", targetingRuleController.updateTargetingRule)
	router.POST("", targetingRuleController.creatingTargetingRule)
}

func (controller *TargetingRuleController) getAllTargetingRule(ctx *gin.Context) {
	controller.TargetingRule = &TargetingRule{}
	targetingRuleService := NewTargetingRuleService(controller.Db, controller.TargetingRule)
	res, err := targetingRuleService.GetAllTargetingRule()
	if err != nil {
		ctx.JSON(500, err)
		return
	}
	ctx.JSON(200, res)
}

func (controller *TargetingRuleController) creatingTargetingRule(ctx *gin.Context) {
	controller.TargetingRule = &TargetingRule{}
	ctx.BindJSON(controller.TargetingRule)
	targetingRuleService := NewTargetingRuleService(controller.Db, controller.TargetingRule)
	res, err := targetingRuleService.CreateTargetingRule()
	if err != nil {
		ctx.JSON(500, err)
		return
	}
	ctx.JSON(200, res)
}

func (controller *TargetingRuleController) updateTargetingRule(ctx *gin.Context) {
	controller.TargetingRule = &TargetingRule{}
	ctx.BindJSON(controller.TargetingRule)
	targetingRuleService := NewTargetingRuleService(controller.Db, controller.TargetingRule)
	res, err := targetingRuleService.UpdateTargetingRule()
	if err != nil {
		ctx.JSON(500, err)
		return
	}
	ctx.JSON(200, res)
}

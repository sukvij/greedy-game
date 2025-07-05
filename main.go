package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sukvij/greedy-game/database"
	servicediscovery "github.com/sukvij/greedy-game/service-discovery"
	targetingrule "github.com/sukvij/greedy-game/targeting-rule"
)

func main() {
	db, err := database.Connection()
	if err != nil {
		return
	}

	db.AutoMigrate(&targetingrule.TargetingRule{})
	fmt.Println(db, err)
	app := gin.Default()
	servicediscovery.RouteService(app, db)
	app.Run(":8080")
}

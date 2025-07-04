package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sukvij/greedy-game/database"
	servicediscovery "github.com/sukvij/greedy-game/service-discovery"
)

func main() {
	db, err := database.Connection()
	if err != nil {
		return
	}
	app := gin.Default()
	servicediscovery.RouteService(app, db)
	app.Run(":8080")

}

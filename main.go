package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sukvij/greedy-game/database"
	"github.com/sukvij/greedy-game/profiling"
	redisservice "github.com/sukvij/greedy-game/redis-service"
	servicediscovery "github.com/sukvij/greedy-game/service-discovery"
	targetingrule "github.com/sukvij/greedy-game/targeting-rule"
)

type User struct {
	Id   int
	Name string
}

func main() {
	db, err := database.Connection()
	if err != nil {
		return
	}

	db.AutoMigrate(&targetingrule.TargetingRule{})
	fmt.Println(db, err)
	redisClient := redisservice.NewRedisClient()
	app := gin.Default()
	profiling.Profiling(app)
	servicediscovery.RouteService(app, db, redisClient)
	app.Run(":8080")

}

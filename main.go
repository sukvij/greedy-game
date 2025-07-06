package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sukvij/greedy-game/database"
	"github.com/sukvij/greedy-game/logs"
	"github.com/sukvij/greedy-game/profiling"
	redisservice "github.com/sukvij/greedy-game/redis-service"
	servicediscovery "github.com/sukvij/greedy-game/service-discovery"
	"github.com/sukvij/greedy-game/tracing"
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

	// db.AutoMigrate(&targetingrule.TargetingRule{})
	// fmt.Println(db, err)
	redisClient := redisservice.NewRedisClient()
	logs := logs.NewAgreeGateLogger()
	tracker := tracing.InitTracer()
	app := gin.Default()
	profiling.Profiling(app)
	servicediscovery.RouteService(app, db, redisClient, logs, tracker)
	app.Run(":8080")

}

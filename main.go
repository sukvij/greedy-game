package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sukvij/greedy-game/gredfers/database"
	"github.com/sukvij/greedy-game/gredfers/logs"
	"github.com/sukvij/greedy-game/gredfers/profiling"
	redisservice "github.com/sukvij/greedy-game/gredfers/redis-service"
	"github.com/sukvij/greedy-game/gredfers/tracing"
	servicediscovery "github.com/sukvij/greedy-game/service-discovery"
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

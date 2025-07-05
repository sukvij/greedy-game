package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sukvij/greedy-game/database"
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
	servicediscovery.RouteService(app, db, redisClient)
	app.Run(":8080")

	// fmt.Println(redisClient)
	// // user := User{Id: 1, Name: "sukariua"}
	// fmt.Println(redis.GetValue(redisClient, "v-1"))
}

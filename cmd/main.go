package main

import (
	"simple-go/configs"
	"simple-go/internal/routes"
	"simple-go/pkg/database"
	"simple-go/pkg/logger"
	"simple-go/pkg/mqtt_client"
	"simple-go/pkg/redis"

	"github.com/gin-gonic/gin"
)

func main() {
	logger.Init()

	configs.InitConfig()

	database.InitDB()

	redis.InitRedis()

	mqtt_client.InitMqtt()

	r := gin.Default()

	routes.RegisterRoutes(r)

	err := r.Run(configs.AppConfig.Port)
	if err != nil {
		return
	}
}

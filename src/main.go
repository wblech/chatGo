package main

import (
	"chatGo/src/application/consumer"
	"chatGo/src/application/controller"
	"chatGo/src/application/controller/middleware"
	"chatGo/src/infrastructure/database/sql"
	"chatGo/src/infrastructure/keycloak"
	"chatGo/src/infrastructure/queue"
	"chatGo/src/infrastructure/settings"
	"github.com/gin-gonic/gin"
)

func main() {
	globalConfig := settings.NewGlobalConfig()
	keycloak.Start(globalConfig)
	db := sql.Start(globalConfig)

	rabbitMQ := queue.NewRabbitMQ(globalConfig)
	defer rabbitMQ.Close()
	qBroker := queue.NewQueueBroker(rabbitMQ)

	consumer.RunAllConsumers(qBroker)

	router := gin.Default()
	routerWithToken := router.Group("/")
	routerWithToken.Use(middleware.ValidateToken)
	controller.Chat(routerWithToken, db)
	controller.WebSocket(router, db, qBroker)
	controller.Auth(router)
	controller.Ping(router)
	controller.Message(router, db)
	controller.Queue(router, qBroker)

	err := router.Run(":8081")
	if err != nil {
		return
	}
}

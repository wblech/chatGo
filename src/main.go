package main

import (
	"chatGo/src/application/consumer"
	"chatGo/src/application/controller"
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
	controller.RouterManager(router, db, qBroker)

	err := router.Run(":8081")
	if err != nil {
		return
	}
}

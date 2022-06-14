package main

import (
	"chatGo/src/application/consumer"
	"chatGo/src/application/controller"
	"chatGo/src/domain/message/repositoryMessage"
	"chatGo/src/infrastructure/database/sql"
	"chatGo/src/infrastructure/keycloak"
	"chatGo/src/infrastructure/queue"
	"chatGo/src/infrastructure/settings"
	"github.com/gin-gonic/gin"
)

func main() {
	globalConfig := settings.NewGlobalConfig()
	keycloak.Start(globalConfig)

	gormSQL := sql.Start(globalConfig)
	db := repositoryMessage.NewRepository(gormSQL)

	rabbitMQ := queue.NewRabbitMQ(globalConfig)
	defer rabbitMQ.Close()
	qBroker := queue.NewQueueBroker(rabbitMQ)

	consumer.RunAllConsumers(qBroker)

	router := gin.Default()
	controller.RouterManager(router, &db, qBroker)

	err := router.Run(":8081")
	if err != nil {
		return
	}
}

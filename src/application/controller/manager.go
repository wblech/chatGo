package controller

import (
	"chatGo/src/application/controller/middleware"
	"chatGo/src/domain/message/repositoryMessage"
	"chatGo/src/infrastructure/queue"
	"github.com/gin-gonic/gin"
)

func RouterManager(router *gin.Engine, db *repositoryMessage.Database, qBroker *queue.Broker) {
	routerWithToken := router.Group("/")
	routerWithToken.Use(middleware.ValidateToken)
	chat(routerWithToken, db)
	websocket(router, db, qBroker)
	auth(router)
	ping(router)
	message(router, db)
	queueController(router, qBroker)
}

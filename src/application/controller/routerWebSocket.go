package controller

import (
	"chatGo/src/domain/message/repositoryMessage"
	"chatGo/src/infrastructure/queue"
	"chatGo/src/infrastructure/socket"
	"github.com/gin-gonic/gin"
)

func websocket(router *gin.Engine, db *repositoryMessage.Database, qBroker *queue.Broker) {
	router.GET("/socket/ws", websocketController(db, qBroker))
}

func websocketController(db *repositoryMessage.Database, qBroker *queue.Broker) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		socket.Execute(c, db, qBroker)
	}

	return fn
}

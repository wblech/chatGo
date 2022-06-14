package controller

import (
	"chatGo/src/domain/message/repository"
	"chatGo/src/infrastructure/queue"
	"chatGo/src/infrastructure/socket"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func websocket(router *gin.Engine, db *gorm.DB, qBroker *queue.Broker) {
	router.GET("/socket/ws", websocketController(db, qBroker))
}

func websocketController(db *gorm.DB, qBroker *queue.Broker) gin.HandlerFunc {
	repo := repository.NewRepository(db)
	fn := func(c *gin.Context) {
		socket.Execute(c, repo, qBroker)
	}

	return fn
}

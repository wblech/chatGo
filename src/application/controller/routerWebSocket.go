package controller

import (
	"chatGo/src/domain/message/repository"
	"chatGo/src/infrastructure/socket"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func WebSocket(router *gin.Engine, db *gorm.DB) {
	router.GET("/socket/ws", websocketController(db))
}

func websocketController(db *gorm.DB) gin.HandlerFunc {
	repo := repository.NewRepository(db)
	fn := func(c *gin.Context) {
		socket.Execute(c, repo)
	}

	return fn
}

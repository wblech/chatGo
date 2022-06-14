package controller

import (
	"chatGo/src/application/controller/middleware"
	"chatGo/src/infrastructure/queue"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RouterManager(router *gin.Engine, db *gorm.DB, qBroker *queue.Broker) {
	routerWithToken := router.Group("/")
	routerWithToken.Use(middleware.ValidateToken)
	chat(routerWithToken, db)
	websocket(router, db, qBroker)
	auth(router)
	ping(router)
	message(router, db)
	queueController(router, qBroker)
}

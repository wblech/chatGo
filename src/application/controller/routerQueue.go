package controller

import (
	"chatGo/src/infrastructure/queue"
	"github.com/gin-gonic/gin"
)

func Queue(router *gin.Engine, qBroker *queue.Broker) {
	//router.Static("/chat", "./public")
	router.POST("/queue", publishMsg(qBroker))
}

type NewMsg struct {
	Msg string `json:"test"`
}

func publishMsg(qBroker *queue.Broker) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		msg := c.Query("msg")
		_ = qBroker.PublishMessage("golang-queue", msg)
		c.JSON(200, "Awesome")
	}
	return fn
}

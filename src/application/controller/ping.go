package controller

import "github.com/gin-gonic/gin"

func Ping(router *gin.Engine) {
	router.GET("/ping", func(c *gin.Context) {
		c.Writer.Write([]byte("pong"))
	})
}

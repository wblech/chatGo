package controller

import "github.com/gin-gonic/gin"

func ping(router *gin.Engine) {
	router.GET("/ping", func(c *gin.Context) {
		c.Writer.Write([]byte("pong"))
	})
}

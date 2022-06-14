package controller

import (
	"chatGo/src/domain/message/read"
	"chatGo/src/domain/message/repositoryMessage"
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

func message(router *gin.Engine, db *repositoryMessage.Database) {
	router.GET("/message", findMessage(db))
}

func findMessage(db *repositoryMessage.Database) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		limit := c.Query("limit")
		if limit == "" {
			_ = c.AbortWithError(404, errors.New("need a limit"))
		}
		limitInt, _ := strconv.Atoi(limit)
		service := read.NewService(db)
		result := service.Execute(limitInt)
		c.JSON(200, result)
	}
	return fn

}

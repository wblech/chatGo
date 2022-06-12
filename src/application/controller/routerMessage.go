package controller

import (
	"chatGo/src/domain/message/repository"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

func Message(router *gin.Engine, db *gorm.DB) {
	//router.Static("/chat", "./public")
	router.GET("/message", findMessage(db))
}

func findMessage(db *gorm.DB) gin.HandlerFunc {
	repo := repository.NewRepository(db)
	fn := func(c *gin.Context) {
		limit := c.Query("limit")
		if limit == "" {
			_ = c.AbortWithError(404, errors.New("need a limit"))
		}
		limitInt, _ := strconv.Atoi(limit)
		result := repo.FindWithLimit(limitInt)
		c.JSON(200, result)
	}
	return fn

}

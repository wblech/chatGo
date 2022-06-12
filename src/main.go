package main

import (
	"chatGo/src/application/controller"
	"chatGo/src/application/controller/middleware"
	"chatGo/src/infrastructure/database/sql"
	"chatGo/src/infrastructure/keycloak"
	"chatGo/src/infrastructure/settings"
	"github.com/gin-gonic/gin"
)

func main() {
	globalConfig := settings.NewGlobalConfig()
	keycloak.Start(globalConfig)
	db := sql.Start(globalConfig)

	router := gin.Default()

	routerWithToken := router.Group("/")
	routerWithToken.Use(middleware.ValidateToken)
	controller.Chat(routerWithToken, db)
	controller.WebSocket(router, db)
	controller.Auth(router)
	controller.Ping(router)
	controller.Message(router, db)

	err := router.Run(":8081")
	if err != nil {
		return
	}
}

package controller

import (
	"chatGo/src/domain/message/repositoryMessage"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func chat(router *gin.RouterGroup, db *repositoryMessage.Database) {
	router.GET("/auth/callback", chatRouter)
}

func chatRouter(c *gin.Context) {
	username, _ := c.Get("preferred_username")
	setCallbackCookie(c.Writer, c.Request, "username", username.(string), false)
	content, err := ioutil.ReadFile("./public/index.html")
	if err != nil {
		http.Error(c.Writer, "Could not open requested file", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(c.Writer, "%s", content)
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)
}

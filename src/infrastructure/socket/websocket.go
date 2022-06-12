package socket

import (
	"chatGo/src/domain/message"
	"chatGo/src/domain/message/repository"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/novalagung/gubrak/v2"
	"log"
	"net/http"
	"strings"
)

const messageNewUser = "New User"
const messageChat = "Chat"
const messageLeave = "Leave"

const maxMessageSize = 1024 * 1024 // 1kb

var connections = make([]*webSocketConnection, 0)

type payload struct {
	Message string
}

type response struct {
	From    string
	Type    string
	Message string
}

type webSocketConnection struct {
	*websocket.Conn
	Username string
}

func Execute(c *gin.Context, db repository.GormDB) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  maxMessageSize,
		WriteBufferSize: maxMessageSize,
	}

	currentGorillaConn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		http.Error(c.Writer, "Could not open websocket connection", http.StatusBadRequest)
	}
	username := c.Query("username")
	currentConn := webSocketConnection{Conn: currentGorillaConn, Username: username}
	connections = append(connections, &currentConn)

	go handleIO(&currentConn, db)
}

func handleIO(currentConn *webSocketConnection, db repository.GormDB) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("ERROR", fmt.Sprintf("%v", r))
		}
	}()

	messageEntitiy := message.NewMessage(currentConn.Username, messageNewUser, "")
	broadcastMessage(currentConn, messageEntitiy)

	messageStr := fmt.Sprintf("User %s: connected", currentConn.Username)
	messageEntitiy = message.NewMessage(currentConn.Username, messageNewUser, messageStr)
	db.Create(messageEntitiy)
	for {
		payload := payload{}
		err := currentConn.ReadJSON(&payload)
		if err != nil {
			if strings.Contains(err.Error(), "websocket: close") {
				messageEntitiy := message.NewMessage(currentConn.Username, messageLeave, "")
				broadcastMessage(currentConn, messageEntitiy)
				ejectConnection(currentConn)
				messageStr := fmt.Sprintf("User %s: disconnect", currentConn.Username)
				messageEntitiy = message.NewMessage(currentConn.Username, messageLeave, messageStr)
				db.Create(messageEntitiy)
				return
			}

			log.Println("ERROR", err.Error())
			continue
		}

		messageEntitiy := message.NewMessage(currentConn.Username, messageChat, payload.Message)
		broadcastMessage(currentConn, messageEntitiy)
		db.Create(messageEntitiy)
	}
}

func ejectConnection(currentConn *webSocketConnection) {
	filtered := gubrak.From(connections).Reject(func(each *webSocketConnection) bool {
		return each == currentConn
	}).Result()
	connections = filtered.([]*webSocketConnection)
}

func broadcastMessage(currentConn *webSocketConnection, message message.Message) {
	for _, eachConn := range connections {
		if eachConn == currentConn {
			continue
		}

		eachConn.WriteJSON(response{
			From:    currentConn.Username,
			Type:    message.Kind,
			Message: message.Message,
		})
	}
}

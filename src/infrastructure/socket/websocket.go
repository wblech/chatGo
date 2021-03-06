package socket

import (
	"chatGo/src/domain/message"
	"chatGo/src/domain/message/repositoryMessage"
	"chatGo/src/infrastructure/queue"
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

type botMsg struct {
	share string
	quote string
}

type webSocketConnection struct {
	*websocket.Conn
	Username string
	BotMsg   botMsg
}

func Execute(c *gin.Context, db *repositoryMessage.Database, qBroker *queue.Broker) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  maxMessageSize,
		WriteBufferSize: maxMessageSize,
	}

	currentGorillaConn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		http.Error(c.Writer, "Could not open websocket connection", http.StatusBadRequest)
	}
	username := c.Query("username")
	msg := botMsg{}
	if username == "" {
		username = "Bot"
		msg.share = c.Query("share")
		msg.quote = c.Query("quote")
	}
	currentConn := webSocketConnection{Conn: currentGorillaConn, Username: username, BotMsg: msg}
	connections = append(connections, &currentConn)

	go handleIO(&currentConn, db, qBroker)
}

func handleIO(currentConn *webSocketConnection, db *repositoryMessage.Database, qBroker *queue.Broker) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("ERROR", fmt.Sprintf("%v", r))
		}
	}()

	var firstMessage message.Message
	if currentConn.Username == "Bot" {
		mess := fmt.Sprintf("%s quote is %s per share\n", currentConn.BotMsg.share, currentConn.BotMsg.quote)
		firstMessage = message.NewMessage(currentConn.Username, messageChat, mess)
		db.Create(firstMessage)
		broadcastMessage(currentConn, firstMessage)
	} else {
		messageEntitiy := message.NewMessage(currentConn.Username, messageNewUser, "")
		broadcastMessage(currentConn, messageEntitiy)

		messageStr := fmt.Sprintf("User %s: connected", currentConn.Username)
		messageEntitiy = message.NewMessage(currentConn.Username, messageNewUser, messageStr)
		db.Create(messageEntitiy)
	}

	for currentConn.Username != "Bot" {
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

		trimStr := strings.TrimSpace(payload.Message)
		splitStr := strings.Split(trimStr, "=")
		var normalMessage message.Message
		if splitStr[0] == "/stock" {
			_ = qBroker.PublishMessage("bot-send", splitStr[1])
			normalMessage = message.NewMessage(currentConn.Username, messageChat, payload.Message)
		} else {
			normalMessage = message.NewMessage(currentConn.Username, messageChat, payload.Message)
			db.Create(normalMessage)
		}
		broadcastMessage(currentConn, normalMessage)
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

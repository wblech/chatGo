package consumer

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"strings"
)

func SendShareInfo(msg string) string {
	host := "localhost:8081"
	sliceStr := strings.Split(msg, " ")
	test := fmt.Sprintf("share=%s&quote=%s", sliceStr[0], sliceStr[1])
	u := url.URL{Scheme: "ws", Host: host, Path: "/socket/ws", RawQuery: test}
	log.Printf("connecting to %s", u.String())

	c, resp, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err, resp.StatusCode)
	}
	defer c.Close()
	err = c.WriteMessage(websocket.TextMessage, []byte(msg))
	err = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		log.Println("write:", err)
		return ""
	}
	return ""
}

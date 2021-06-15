package v1

import (
	"log"
	"net/http"
	"nmb/pkg"
	"time"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 2048
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	hub *pkg.Hub
)

func init() {
	hub = pkg.NewHub()
	go hub.Run()
}

func ServeWS(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := pkg.NewClient(conn, hub)
	hub.Register(client)

	go client.Read()
	go client.Write()
	startup(client)
}

func startup(c *pkg.Client) {
	for channel := range channels {
		message := pkg.Message{
			Sender:    "server",
			Channel:   channel,
			Kind:      4,
			Timestamp: time.Now(),
			Payload:   nil,
		}
		raw, _ := json.Marshal(message)
		c.Inject(raw)
	}
}

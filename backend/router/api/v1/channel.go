package v1

import (
	"net/http"

	"encoding/json"
	"nmb/pkg"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	channels map[string]bool
)

type Channel struct {
	Name string `json:"channel"`
	Type int    `json:"type"`
}

func init() {
	channels = make(map[string]bool)
}

func CreateChannel(c *gin.Context) {
	channel := &Channel{}
	if err := c.ShouldBindJSON(channel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
	}

	if _, ok := channels[channel.Name]; ok {
		c.JSON(http.StatusOK, gin.H{
			"message": "channel already exists",
		})
	} else {
		channels[channel.Name] = true
		c.JSON(http.StatusOK, gin.H{
			"message": "channel created",
		})
		message := pkg.Message{
			Sender:    "server",
			Channel:   channel.Name,
			Kind:      4,
			Payload:   nil,
			Timestamp: time.Now(),
		}
		raw, _ := json.Marshal(message)
		hub.Publish(raw)
	}
}

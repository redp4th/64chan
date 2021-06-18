package v1

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"nmb/pkg"
	"time"

	"github.com/gin-gonic/gin"
)

func FileHandler(c *gin.Context) {
	file, _ := c.FormFile("file")
	fileURL := "https://10.128.196.86:8000/static/" + file.Filename
	sender := c.PostForm("sender")
	channel := c.PostForm("channel")
	c.SaveUploadedFile(file, "/Users/redpath/work/nmb/storage/"+file.Filename)
	message := pkg.Message{
		Sender:    sender,
		Channel:   channel,
		Kind:      3,
		Payload:   fileURL,
		Timestamp: time.Now(),
	}
	raw, _ := json.Marshal(message)
	hub.Publish(raw)
}

func PicHandler(c *gin.Context) {
	var raw []byte
	pic, _ := c.FormFile("file")
	sender := c.PostForm("sender")
	channel := c.PostForm("channel")
	message := pkg.Message{
		Sender:    sender,
		Channel:   channel,
		Kind:      2,
		Payload:   "Error",
		Timestamp: time.Now(),
	}
	file, err := pic.Open()

	if err != nil {
		raw, _ = json.Marshal(message)
		hub.Publish(raw)
		return
	}
	defer file.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		raw, _ = json.Marshal(message)
		hub.Publish(raw)
		return
	}

	b64str := base64.StdEncoding.EncodeToString(buf.Bytes())
	message.Payload = b64str

	raw, _ = json.Marshal(message)
	hub.Publish(raw)
}

package pkg

import (
	"time"
)

const (
	Text = iota
	Picture
	File
	Control
)

type Message struct {
	Sender    string    `json:"sender"`
	Channel   string    `json:"channel"`
	Kind      uint64    `json:"kind"`
	Timestamp time.Time `json:"timestamp"`
	Payload   string    `json:"payload"`
}

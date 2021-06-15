package pkg

import (
	"time"
)

const (
	text = iota
	picture
	file
	control
)

type Message struct {
	Sender    string    `json:"sender"`
	Channel   string    `json:"channel"`
	Kind      uint64    `json:"kind"`
	Timestamp time.Time `json:"timestamp"`
	Payload   []byte    `json:"payload"`
}

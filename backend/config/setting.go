package config

import (
	"log"

	"github.com/go-ini/ini"
)

var (
	JWTSecret []byte
)

func init() {
	cfg, err := ini.Load("config/config.ini")
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}

	JWTSecret = []byte(cfg.Section("app").Key("SECRET").String())
}

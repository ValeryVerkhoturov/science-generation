package config

import (
	"github.com/joho/godotenv"
	"os"
)

const (
	SessionName    = "chat-widget"
	SessionUserKey = "userId"
)

var (
	Port string
	Host string
)

func init() {
	var err error

	if os.Getenv("GO_ENV") != "production" {
		if err = godotenv.Load(); err != nil {
			panic(err)
		}
	}

	Port = os.Getenv("PORT")
	Host = os.Getenv("HOST")

	if len(Port) == 0 || len(Host) == 0 {
		panic("Invalid env variables")
	}
}

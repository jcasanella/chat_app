package main

import (
	"github.com/jcasanella/chat_app/config"
	"github.com/jcasanella/chat_app/jwt"
	"github.com/jcasanella/chat_app/server"
)

func main() {
	jwt.Init(64)

	config.Init("local")
	server.Init()
}

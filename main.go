package main

import (
	"github.com/jcasanella/chat_app/config"
	"github.com/jcasanella/chat_app/server"
)

func main() {
	config.Init("local")
	server.Init()
}

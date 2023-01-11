package main

import (
	"github.com/jcasanella/chat_app/config"
	"github.com/jcasanella/chat_app/repository"
	"github.com/jcasanella/chat_app/security"
	"github.com/jcasanella/chat_app/server"
)

func main() {
	security.Init(64)
	config.Init("local")

	st := repository.NewMemStorage()
	s := server.NewServer(st)
	s.StartServer()
}

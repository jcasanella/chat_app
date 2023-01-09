package main

import (
	"github.com/jcasanella/chat_app/config"
	"github.com/jcasanella/chat_app/repository"
	"github.com/jcasanella/chat_app/security"
	"github.com/jcasanella/chat_app/server"
)

func main() {
	security.Init(64)

	// Init database
	st := repository.NewMemStorage()
	repository.NewServiceDb(st)

	config.Init("local")
	server.Init()
}

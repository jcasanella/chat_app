package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jcasanella/chat_app/config"
	"github.com/jcasanella/chat_app/controller"
	"github.com/jcasanella/chat_app/repository"
)

func NewRouter(s repository.Storage) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	v1 := router.Group("v1")
	{
		sDb := repository.NewServiceDb(s)
		lc := controller.NewLoginController(*sDb)
		v1.GET("login", lc.Login)
	}

	return router
}

func Init(s repository.Storage) {
	config := config.GetConfig()
	r := NewRouter(s)
	r.Run(fmt.Sprintf(":%s", config.GetString("port")))
}

package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jcasanella/chat_app/config"
	"github.com/jcasanella/chat_app/controller"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	v1 := router.Group("v1")
	{
		lc := new(controller.LoginController)
		v1.GET("login", lc.Login)
	}

	return router
}

func Init() {
	config := config.GetConfig()
	r := NewRouter()
	r.Run(fmt.Sprintf(":%s", config.GetString("port")))
}

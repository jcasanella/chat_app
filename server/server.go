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
		login := v1.Group("login")
		{
			lc := new(controller.LoginController)
			login.GET("/", lc.Status)
		}
	}

	return router
}

func Init() {
	config := config.GetConfig()
	r := NewRouter()
	r.Run(fmt.Sprintf(":%s", config.GetString("port")))
}

package server

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/jcasanella/chat_app/config"
	"github.com/jcasanella/chat_app/controller"
	"github.com/jcasanella/chat_app/repository"
	"github.com/jcasanella/chat_app/service"
)

type Server struct {
	UserService *service.UserService
}

func NewServer(storage repository.Storage) *Server {
	db := repository.NewServiceDb(storage)

	return &Server{
		UserService: service.NewUserService(db),
	}
}

func (s Server) newRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	v1 := router.Group("v1")
	{
		lc := controller.NewLoginController(s.UserService)
		v1.GET("login", lc.Login)
		v1.POST("register", lc.Register)
	}

	return router
}

func (s Server) StartServer() {
	config := config.GetConfig()
	r := s.newRouter()
	r.Run(fmt.Sprintf(":%s", config.GetString("port")))
}

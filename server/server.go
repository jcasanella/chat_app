package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/jcasanella/chat_app/config"
	"github.com/jcasanella/chat_app/controller"
	"github.com/jcasanella/chat_app/repository"
	"github.com/jcasanella/chat_app/service"
)

type Server struct {
	UserService *service.UserService
}

// NewServer receives the Storate object and initialize a Server
func NewServer(storage repository.Storage) *Server {
	db := repository.NewServiceDb(storage)

	return &Server{
		UserService: service.NewUserService(db),
	}
}

func getTemplatesFolder() string {
	templates := "./templates"
	if _, err := os.Stat(templates); os.IsNotExist(err) {
		return "./../templates"
	}

	return templates
}

func (s Server) newRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Static("/static", "./static")
	templatesFolder := getTemplatesFolder()
	router.LoadHTMLGlob(fmt.Sprintf("%s/*.html", templatesFolder))
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{"content": "This is the index page"})
	})

	v1 := router.Group("v1")
	{
		lc := controller.NewLoginController(s.UserService)
		v1.GET("login", lc.Login)
		v1.POST("register", lc.Register)
	}

	return router
}

// StartServer starts the server in the port indicated in the config file
func (s Server) StartServer() {
	config := config.GetConfig()
	r := s.newRouter()
	r.Run(fmt.Sprintf(":%s", config.GetString("port")))
}

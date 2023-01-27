package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jcasanella/chat_app/model"
	"github.com/jcasanella/chat_app/security"
	"github.com/jcasanella/chat_app/service"
)

type LoginController struct {
	uService *service.UserService
}

func NewLoginController(u *service.UserService) *LoginController {
	return &LoginController{
		uService: u,
	}
}

func (lc LoginController) Login(c *gin.Context) {
	l := model.User{}
	if err := c.BindJSON(&l); err != nil {
		printError(c, err)
	} else {
		if _, err := lc.uService.GetUser(l); err != nil {
			printError(c, err)
		}

		if sig, err := security.GenerateJWT(l.Name); err != nil {
			printError(c, err)
		} else {
			c.JSON(http.StatusOK, gin.H{"token": sig})
		}
	}
}

func (lc LoginController) Register(c *gin.Context) {
	u := model.User{}
	if err := c.BindJSON(&u); err != nil {
		printError(c, err)
	} else {
		if _, err := lc.uService.AddUser(u); err != nil {
			printError(c, err)
		} else {
			c.JSON(http.StatusCreated, u)
		}
	}
}

func printError(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	c.Abort()
}

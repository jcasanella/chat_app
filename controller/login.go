package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jcasanella/chat_app/model"
)

type LoginController struct{}

func (lc LoginController) Status(c *gin.Context) {
	var user model.User
	c.Bind(&user)
	if user.Name != "" && user.Password != "" {
		c.String(http.StatusOK, "Valid user")
	} else {
		c.String(http.StatusBadRequest, "Invalid user")
	}
}

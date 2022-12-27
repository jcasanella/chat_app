package controller

import (
	"fmt"
	"net/http"

	"github.com/jcasanella/chat_app/jwt"

	"github.com/gin-gonic/gin"
)

type LoginController struct{}

func (lc LoginController) Login(c *gin.Context) {
	n := c.Query("name")
	p := c.Query("password")
	if n != "" && p != "" {
		fmt.Printf("Token generated: %s\n", jwt.Token)
		c.String(http.StatusOK, "Valid user")
	} else {
		c.String(http.StatusBadRequest, "Invalid user")
	}
}

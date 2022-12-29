package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jcasanella/chat_app/security"
)

type LoginController struct{}

func (lc LoginController) Login(c *gin.Context) {
	n := c.Query("name")
	p := c.Query("password")
	if n != "" && p != "" {
		fmt.Printf("Token generated: %s\n", security.SecretKey)
		sig, _ := security.GenerateJWT(n)
		fmt.Printf("Signature: %s\n", sig)
		c.String(http.StatusOK, "Valid user")
	} else {
		c.String(http.StatusBadRequest, "Invalid user")
	}
}

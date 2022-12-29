package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jcasanella/chat_app/security"
)

type LoginController struct{}

func (lc LoginController) Login(c *gin.Context) {
	n := c.Query("name")
	p := c.Query("password")
	if n != "" && p != "" {
		if sig, err := security.GenerateJWT(n); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
		} else {
			c.JSON(http.StatusOK, gin.H{"token": sig})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user"})
		c.Abort()
	}
}

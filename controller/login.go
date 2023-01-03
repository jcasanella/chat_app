package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jcasanella/chat_app/security"
)

type LoginController struct{}

type UserTest struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (lc LoginController) Login(c *gin.Context) {
	l := UserTest{}
	if err := c.BindJSON(&l); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user"})
		c.Abort()
	} else {
		if sig, err := security.GenerateJWT(l.Name); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
		} else {
			c.JSON(http.StatusOK, gin.H{"token": sig})
		}
	}
}

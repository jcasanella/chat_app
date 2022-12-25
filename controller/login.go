package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginController struct{}

func (lc LoginController) Login(c *gin.Context) {
	n := c.Query("name")
	p := c.Query("password")
	if n != "" && p != "" {
		c.String(http.StatusOK, "Valid user")
	} else {
		c.String(http.StatusBadRequest, "Invalid user")
	}
}

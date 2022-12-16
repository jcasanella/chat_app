package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginController struct{}

func (lc LoginController) Status(c *gin.Context) {
	c.String(http.StatusOK, "working!!!")
}

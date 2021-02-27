package system

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Panel(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ResponseData(code int, msg string, data interface{}) gin.H {
	return gin.H{
		"code": code,
		"msg": msg,
		"data": data,
	}
}

func Response(c *gin.Context, data gin.H) {
	c.JSON(http.StatusOK, data)
}

func ResponseOk(c *gin.Context, data interface{}) {
	Response(c, ResponseData(0, "", data))
}

func ResponseErr(c *gin.Context, code int, msg string, data interface{}) {
	Response(c, ResponseData(code, msg, data))
}
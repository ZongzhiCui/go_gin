package common

import (
	"time"

	"github.com/gin-gonic/gin"
)

func OutputJson(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(200, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
		"time": time.Now(),
	})
}

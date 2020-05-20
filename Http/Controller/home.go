package Controller

import (
	"github.com/ZongzhiCui/go_gin/common"

	"github.com/gin-gonic/gin"
)

func LoginEndpoint(c *gin.Context) {
	name := c.PostForm("name")
	age := c.DefaultQuery("age", "15")

	data := map[string]interface{}{
		"param1": name,
		"param2": age,
	}
	common.OutputJson(c, 200, "OK", data)
}

func Ping(c *gin.Context) {
	common.OutputJson(c, 200, "pong", []int{1, 2})
}

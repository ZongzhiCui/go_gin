package main

import (
	"github.com/ZongzhiCui/go_gin/Http/Controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	v1 := r.Group("/api") //路由组,可嵌套
	{
		//解决 跨域问题 ->
		//var c *gin.Context
		//c.Header("Access-Control-Allow-Origin", "*")
		//c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		//c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		//c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		//c.Header("Access-Control-Allow-Credentials", "true")
		//
		////放行所有OPTIONS方法
		//method := c.Request.Method
		//if method == "OPTIONS" {
		//	c.AbortWithStatus(http.StatusNoContent)
		//}
		//解决 跨域问题 <-

		v1.POST("/login", Controller.LoginEndpoint)
		v1.POST("/create", Controller.Create_user)

		v1.GET("/ping", Controller.Ping)

	}

	r.Run(":3001") // 监听并在 0.0.0.0:8080 上启动服务
	//然而 127.0.0.1:3001 可以访问, 0.0.0.0:3001 却不能访问
}

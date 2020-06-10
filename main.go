package main

import (
	"io"
	"os"
	"time"

	"ZongzhiCui/go_gin/Http/Controller"

	"github.com/gin-gonic/gin"
)

var basePath = "storage/log/"

func main() {
	gin.DisableConsoleColor()
	//dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	_ = os.MkdirAll(basePath, os.ModePerm)
	prefix := "gin_" + time.Now().Format("20060102")

	infoFile, _ := os.Create(basePath + prefix + "_info.log")
	gin.DefaultWriter = io.MultiWriter(infoFile, os.Stdout)

	errorFile, _ := os.Create(basePath + prefix + "_error.log")
	gin.DefaultErrorWriter = io.MultiWriter(errorFile, os.Stdout)

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

	//优雅地重启或停止 不适用Windows;Windows不提供复杂的过程控制或信号,会提示undefined: syscall.SIGUSR1
	//信号: 异步的通知机制,类似事件
	/* 查看所有信号: kill -l
	ctrl + c: SIGINT 强制进程结束
	ctrl + z: SIGTSTP 任务中断，进程挂起
	ctrl + \: SIGQUIT 进程结束 和 dump core
	ctrl + d: EOF*/

	//endless.DefaultReadTimeOut time.Duration
	//endless.DefaultWriteTimeOut time.Duration
	//endless.DefaultMaxHeaderBytes = 1 << 20
	//_ = endless.ListenAndServe(":3001", r)

	//server := endless.NewServer("localhost:3001", r)
	//server.BeforeBegin = func(add string) {
	//	log.Printf("Actual pid is %d", syscall.Getpid())
	//	// save it somehow
	//}
	//_ = server.ListenAndServe()
}

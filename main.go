package main

import (
	"github.com/gin-gonic/gin"
	"ticket/middleware"
	"ticket/routers"
	"ticket/util"
)

func main() {
	r := gin.Default()
	r.Use(middleware.LoggerMiddleware())
	r.Use(util.Cors())

	r.GET("/", routers.Hello)
	r.POST("/checkUser", routers.Login)
	//用户注册
	r.POST("/addUser", routers.Register)

	r.POST("/user/upload", routers.Upload)
	_ = r.Run(":8096")
}

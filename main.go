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

	//上传小票
	r.POST("/user/upload", routers.Upload)

	//保存小票
	r.POST("/user/saveTicket", routers.SaveTicket)
	_ = r.Run(":8096")
}

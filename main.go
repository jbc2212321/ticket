package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ticket/routers"
	"ticket/util"
)

func main() {
	r := gin.Default()
	r.Use(util.Cors())

	r.GET("/",routers.Hello )
	r.POST("/checkUser", routers.Login)

	_ = r.Run(":8096")
	fmt.Println("???")
}

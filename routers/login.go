package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ticket/util"
)

type LoginParam struct {
	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
	User     string `form:"username" json:"phone" uri:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var json LoginParam
	// 将request的body中的数据，自动按照json格式解析到结构体
	if err := c.ShouldBindJSON(&json); err != nil {
		// 返回错误信息
		// gin.H封装了生成json数据的工具
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 判断用户名密码是否正确
	if json.User != "root" || json.Password != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
		return
	}
	resp := util.Response{}
	resp["status"] = "200"
	resp["message"] = "OK"
	c.JSON(http.StatusOK, resp)
}

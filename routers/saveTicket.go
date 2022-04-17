package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"ticket/database"
	"ticket/middleware"
	"ticket/util"
)

type SaveTicketParam struct {
	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Category string `json:"value" binding:"required"`
}

//上传图片
func SaveTicket(c *gin.Context) {
	var json database.VatInvoice
	resp := util.GetResponse()

	fmt.Println("json:", json)
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		middleware.Log.Infof("json解析失败[%s]", c.Request)
		fmt.Println("err:", err)
		return
	}
	fmt.Println("*****************:", json)

	resp.Message = "保存成功"
	c.JSON(http.StatusOK, resp)
}

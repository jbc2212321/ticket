package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ticket/util"
)

type UploadParam struct {
	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Category string `json:"value" binding:"required"`
}

//上传图片
func Upload(c *gin.Context) {
	//var json UploadParam
	resp := util.GetResponse()

	f, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, "接收文件失败")
		return
	}
	if err := c.SaveUploadedFile(f, f.Filename); err != nil {
		c.String(http.StatusBadRequest, "保存文件失败")
		return
	}
	c.String(http.StatusOK, "上传文件成功")

	resp.Message = "注册成功"
	c.JSON(http.StatusOK, resp)
}

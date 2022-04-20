package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ticket/database"
	"ticket/middleware"
	"ticket/util"
)

var imageDao database.ImageImpl

type SaveTicketParam struct {
	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
	//Username string `json:"username" binding:"required"`
	//Password string `json:"password" binding:"required"`
	ImageName  string              `json:"imageName" binding:"required"`
	TicketImg  string              `json:"ticketImg" binding:"required"`
	VatInvoice database.VatInvoice `json:"vatInvoice" binding:"required"`
}

//保存识别后的小票
func SaveTicket(c *gin.Context) {
	//var json database.VatInvoice
	var json SaveTicketParam
	resp := util.GetResponse()

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		middleware.Log.Infof("json解析失败[%s]", c.Request)
		//fmt.Println("err:", err)
		resp.Status = util.JSONError
		return
	}
	//imgPath := util.ImagePath + json.ImageName

	//img:=database.Image{
	//	Id:         util.GetSnowflakeId(),
	//	TicketId: util.GetSnowflakeId(),
	//	BinaryData: filebytes,
	//	OcrBinaryData: ,
	//	Type:       0,
	//	CreateTime: time.Now(),
	//}
	//imageDao.AddImage()
	resp.Status = util.SUCCESS
	resp.Message = "保存成功"
	c.JSON(http.StatusOK, resp)
}

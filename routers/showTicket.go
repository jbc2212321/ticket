package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"ticket/middleware"
	"ticket/util"
	"time"
)

type ShowTicketParam struct {
	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
	//Username string `json:"username" binding:"required"`
	//Password string `json:"password" binding:"required"`
	//UserId     string              `json:"userId" binding:"required"`
	TicketId string `json:"ticketId" binding:"required"`
}

type ListTicketByUserIdParam struct {
	UserId string `json:"userId" binding:"required"`
}

//展示小票
func ShowTicket(c *gin.Context) {
	//var json database.VatInvoice
	var json ShowTicketParam
	resp := util.GetResponse()

	if err := c.ShouldBindJSON(&json); err != nil {
		fmt.Println("解析失败！")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		middleware.Log.Infof("json解析失败[%s]", c.Request)
		resp.Status = util.JSONError
		return
	}

	images, err := imageDao.GetImageByTicketId(util.TranToInt64(json.TicketId))
	if err != nil {
		resp.Status = util.DBError
		return
	}
	resp.Status = util.SUCCESS
	resp.Message = "读取成功"
	type Image struct {
		OcrBinaryData string
		BinaryData    string
	}
	img := Image{
		string(images[0].OcrBinaryData),
		string(images[0].BinaryData),
	}
	resp.Data = img
	//fmt.Println(string(images[0].BinaryData))
	c.JSON(http.StatusOK, resp)
}

//小票列表
func ListTicketByUserId(c *gin.Context) {
	//var json database.VatInvoice
	var json ListTicketByUserIdParam
	resp := util.GetResponse()

	if err := c.ShouldBindJSON(&json); err != nil {
		fmt.Println("解析失败！")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		middleware.Log.Infof("json解析失败[%s]", c.Request)
		resp.Status = util.JSONError
		return
	}

	vats, err := vatDao.GetVatByUserId(util.TranToInt64(json.UserId))
	if err != nil {
		resp.Status = util.DBError
		return
	}
	ticketIds := make([]int64, len(vats))

	for _, i := range vats {
		ticketIds = append(ticketIds, i.Id)
	}

	//小票列表
	images, err := imageDao.GetImageByTicketId(ticketIds...)
	if err != nil {
		resp.Status = util.DBError
		return
	}
	resp.Status = util.SUCCESS
	resp.Message = "读取成功"

	type ImageList struct {
		TicketId string    `json:"ticket_id"`
		Time     time.Time `json:"create_time"`
		Category string    `json:"type"`
	}
	imageList := make([]ImageList, 0)
	for _, v := range images {
		image := ImageList{
			TicketId: util.TranToString(v.TicketId),
			Time:     v.CreateTime,
		}
		if v.Type == 0 {
			image.Category = "增值税发票"
		}
		imageList = append(imageList, image)
	}
	resp.Data = imageList
	//fmt.Println(images[0].TicketId)
	c.JSON(http.StatusOK, resp)
}

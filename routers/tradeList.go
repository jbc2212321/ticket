package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ticket/database"
	"ticket/util"
)

//import (
//	//	"fmt"
//	"github.com/gin-gonic/gin"
//	"net/http"
//	"ticket/database"
//	"ticket/util"
//)

type ShowParam struct {
	Num      string `json:"num" binding:"required"`
	Songid   string `json:"songid" binding:"required"`
	Songname string `json:"songname" binding:"required"`
	Owner    string `json:"owner" binding:"required"`
	Ownerid  string `json:"ownerid" binding:"required"`
	Buyerid  string `json:"buyerid" binding:"required"`
	Status   string `json:"status" binding:"required"`
}

var tradeListDao database.TradelistImpl

func GetTradeList(c *gin.Context) {
	resp := util.GetResponse()

	resp.Status = util.SUCCESS
	resp.Message = "读取成功"
	//fmt.Println(json.status)

	tradeLists, _ := tradeListDao.GetTradeList()
	resp.Data = tradeLists
	//verifyList, _ := verifyListDao.GetVerifyByStatus(util.TranToInt64(json.Status))

	c.JSON(http.StatusOK, resp)
}

func Buy(c *gin.Context) {

	resp := util.GetResponse()

	resp.Status = util.SUCCESS
	resp.Message = "读取成功"
	//fmt.Println(json.status)

	//tradeLists,_:=tradeListDao.GetVerifyList();
	//resp.Data=tradeLists
	//verifyList, _ := verifyListDao.GetVerifyByStatus(util.TranToInt64(json.Status))

	c.JSON(http.StatusOK, resp)
}

type AddTradeListParam struct {
	UserId  string `json:"ownerid" binding:"required"`
	BuyerId string `json:"buyerid" binding:"required"`
	SongId  string `json:"songid" binding:"required"`
	Status  string `json:"status" binding:"required"`
}

////保存到trainlist 提交交易
//func AddTradeList(c *gin.Context) {
//	//var json database.VatInvoice
//	var json AddTradeListParam
//	resp := util.GetResponse()
//
//	if err := c.ShouldBindJSON(&json); err != nil {
//		fmt.Println("解析失败！")
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		middleware.Log.Infof("json解析失败[%s]", c.Request)
//		//fmt.Println("err:", err)
//		resp.Status = util.JSONError
//		return
//	}
//
//	usernames, _ := userDao.GetUserById(util.TranToInt64(json.UserId), util.TranToInt64(json.BuyerId))
//	//_, _ = songImpl.GetSong()
//	tradelist := &database.Tradelist{
//		Songid:   util.TranToInt(json.SongId),
//		Songname: json.SongName,
//		Owner:    usernames[0].Username,
//		Buyer:    usernames[1].Username,
//		Ownerid:  util.TranToInt64(json.UserId),
//		Buyerid:  util.TranToInt64(json.BuyerId),
//		Status:   util.TranToInt(json.Status),
//	}
//	_ = tradeListDao.AddTradeList()
//	err := vatDao.AddVat(vat)
//	if err != nil {
//		resp.Status = util.DBError
//		return
//	}
//
//	//保存到image表中
//	imgPath := util.ImagePath + json.ImageName
//	fileBytes, err := ioutil.ReadFile(imgPath)
//	if err != nil {
//		fmt.Println(err)
//	}
//	ticketImage := "data:image/png;base64," + base64.StdEncoding.EncodeToString(fileBytes)
//
//	img := &database.Image{
//		Id:            util.GetSnowflakeId(),
//		TicketId:      vat.Id,
//		BinaryData:    []byte(ticketImage),
//		OcrBinaryData: []byte(json.TicketImg),
//		Type:          0,
//		CreateTime:    time.Now(),
//	}
//
//	err = imageDao.AddImage(img)
//	if err != nil {
//		resp.Status = util.DBError
//		return
//	}
//
//	_ = logDao.AddLog(GetLog(util.TranToInt64(json.UserId), "保存发票"))
//
//	resp.Status = util.SUCCESS
//	resp.Message = "保存成功"
//	c.JSON(http.StatusOK, resp)
//}

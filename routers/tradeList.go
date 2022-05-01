package routers

import (
	//	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"ticket/database"
	"ticket/util"
)

type BuyParam struct {
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

var belongDao database.BelongImpl
var songDao database.SongsImpl

func Buy(c *gin.Context) {
	var json BuyParam

	resp := util.GetResponse()
	if err := c.ShouldBindJSON(&json); err != nil {
		fmt.Println("解析失败！")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		middleware.Log.Infof("json解析失败[%s]", c.Request)
		resp.Status = util.JSONError
		return
	}

	resp.Status = util.SUCCESS
	resp.Message = "读取成功"
	err := tradeListDao.DelTradeListById(util.TranToInt64(json.Num)) //imageDao.AddImage(img)
	if err != nil {
		resp.Status = util.DBError
		return
	}
	userName, err := userDao.GetUserById(util.TranToInt64(json.Buyerid))
	fmt.Println("buyerName:", userName[0].Username)

	songs, err := songDao.GetSongByID(util.TranToInt64(json.Songid))
	fmt.Println("songName:", songs[0].Name)
	if err != nil {
		resp.Status = util.DBError
		return
	}
	if json.Status == "1" {
		fmt.Println("json status:", 1)
		belong := &database.Belong{
			Userid:   util.TranToInt64(json.Buyerid),
			Username: userName[0].Username,
			Songid:   int(util.TranToInt64(json.Songid)),
			Songname: songs[0].Name,
		}
		err = belongDao.AddBelong(belong)
	} else {
		err = belongDao.ChangeBelongBySongID(util.TranToInt64(json.Songid), userName[0].Username, util.TranToInt64(json.Buyerid))
	}
	if err != nil {
		resp.Status = util.DBError
		return
	}

	if json.Status == "1" {
		verifyList := &database.Verifylist{
			Userid:   util.TranToInt64(json.Buyerid),
			Username: userName[0].Username,
			Songid:   int(util.TranToInt64(json.Songid)),
			Songname: songs[0].Name,
			Status:   1,
		}
		err = verifyListDao.AddVerifyList(verifyList)
	} else {
		err = verifyListDao.ChangeVerifyListBySongID(util.TranToInt64(json.Songid), userName[0].Username, util.TranToInt64(json.Buyerid))
		//	err=belongDao.ChangeBelongBySongID(util.TranToInt64(json.Songid),userName[0].Username,util.TranToInt64(json.Buyerid))
	}
	if err != nil {
		resp.Status = util.DBError
		return
	}
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

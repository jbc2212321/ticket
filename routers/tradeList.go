package routers

import (
	"fmt"
	"ticket/middleware"
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

package routers

import (
	//	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"ticket/database"
	"ticket/util"
)

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

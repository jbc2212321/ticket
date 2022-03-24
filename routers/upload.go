package routers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"ticket/database"
	"ticket/middleware"
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
		resp.Message = "接收文件失败"
		resp.Code = util.IOError
		c.JSON(http.StatusOK, resp)
		return
	}
	//将读取到的文件保存到本地(服务端)
	dst := fmt.Sprintf("image/%s", f.Filename)
	fmt.Println("dst:", dst)
	if err := c.SaveUploadedFile(f, dst); err != nil {
		resp.Message = "保存文件失败"
		resp.Code = util.IOError
		c.JSON(http.StatusOK, resp)
		return
	}

	v, err := OCRTicket(dst)
	if err != nil {
		resp.Message = "解析文件失败"
		resp.Code = util.IOError
		c.JSON(http.StatusOK, resp)
		return
	}

	VatTicket := TranToVatDao(v)
	resp.Data = VatTicket
	fmt.Println("小写:", VatTicket.AmountInFiguers)
	resp.Code = util.SUCCESS
	resp.Message = "上传文件成功"
	c.JSON(http.StatusOK, resp)
}

func OCRTicket(dst string) (*util.VatInvoice, error) {
	v := &util.VatInvoice{}
	var host = "https://aip.baidubce.com/rest/2.0/ocr/v1/vat_invoice"
	uri, err := url.Parse(host)
	if err != nil {
		fmt.Println(err)
	}
	query := uri.Query()
	query.Set("access_token", util.AccessToken)
	uri.RawQuery = query.Encode()

	filebytes, err := ioutil.ReadFile(dst)
	if err != nil {
		fmt.Println(err)
	}

	image := base64.StdEncoding.EncodeToString(filebytes)
	sendBody := http.Request{}
	sendBody.ParseForm()
	sendBody.Form.Add("image", image)
	sendData := sendBody.Form.Encode()

	client := &http.Client{}
	request, err := http.NewRequest("POST", uri.String(), strings.NewReader(sendData))
	if err != nil {
		fmt.Println(err)
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := client.Do(request)
	defer response.Body.Close()
	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(result))

	//读取的数据为json格式，需要进行解码
	err = json.Unmarshal(result, &v)
	if err != nil {
		middleware.Log.Infof("解码失败,文件:[%+v]", v)
	}
	return v, nil
}

func readTicket(dst string) error {
	data, err := ioutil.ReadFile(dst)
	if err != nil {
		middleware.Log.Infof("读取文件失败,文件:[%s]", dst)
		return err
	}

	v := &util.VatInvoice{}
	//读取的数据为json格式，需要进行解码
	err = json.Unmarshal(data, &v)
	if err != nil {
		middleware.Log.Infof("解码失败,文件:[%+v]", v)
		return err
	}

	fmt.Println(util.TranToStringFromStruct(v.WordsResult.CommodityName))
	//AddTicket(v)
	//fmt.Println(string(v.WordsResult.CommodityAmount))
	//fmt.Println(v.WordsResult.CommodityName)
	//fmt.Println(v.WordsResult.CommodityTax)
	//fmt.Println(v.WordsResult.Password)    //密码区
	//fmt.Println(v.WordsResult.InvoiceType) //发票种类
	return nil
}

//VatInvoice to database VatInvoice
func TranToVatDao(v *util.VatInvoice) *database.VatInvoice {
	vat := &database.VatInvoice{}
	vat.InvoiceCode = v.WordsResult.InvoiceCode
	vat.MachineCode = v.WordsResult.MachineCode
	vat.InvoiceTypeOrg = v.WordsResult.InvoiceTypeOrg
	vat.InvoiceNum = v.WordsResult.InvoiceNum
	vat.InvoiceDate = v.WordsResult.InvoiceDate
	vat.PurchaserName = v.WordsResult.PurchaserName
	vat.PurchaserRegisterNum = v.WordsResult.PurchaserRegisterNum
	vat.PurchaserAddress = v.WordsResult.PurchaserAddress
	vat.PurchaserBank = v.WordsResult.PurchaserBank
	vat.Password = v.WordsResult.Password
	vat.CommodityName = util.TranToStringFromStruct(v.WordsResult.CommodityName)
	vat.CommodityType = util.TranToStringFromStruct(v.WordsResult.CommodityType)
	vat.CommodityUnit = util.TranToStringFromStruct(v.WordsResult.CommodityUnit)
	vat.CommodityNum = util.TranToStringFromStruct(v.WordsResult.CommodityNum)
	vat.CommodityPrice = util.TranToStringFromStruct(v.WordsResult.CommodityPrice)
	vat.CommodityAmount = util.TranToStringFromStruct(v.WordsResult.CommodityAmount)
	vat.CommodityTaxRate = util.TranToStringFromStruct(v.WordsResult.CommodityTaxRate)
	vat.CommodityTax = util.TranToStringFromStruct(v.WordsResult.CommodityTax)
	vat.TotalAmount = v.WordsResult.TotalAmount
	vat.TotalTax = v.WordsResult.TotalTax
	vat.AmountInWords = v.WordsResult.AmountInWords
	vat.AmountInFiguers = v.WordsResult.AmountInFiguers
	vat.SellerName = v.WordsResult.SellerName
	vat.SellerRegisterNum = v.WordsResult.SellerRegisterNum
	vat.SellerAddress = v.WordsResult.SellerAddress
	vat.SellerBank = v.WordsResult.SellerBank
	vat.Remarks = v.WordsResult.Remarks
	vat.Payee = v.WordsResult.Payee
	vat.Checker = v.WordsResult.Checker
	vat.NoteDrawer = v.WordsResult.NoteDrawer
	return vat
}

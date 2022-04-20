package test

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io/ioutil"
	"testing"
	"ticket/database"
	"ticket/util"
	"time"
)

var userDao database.UserImpl
var imageDao database.ImageImpl

func TestDb(t *testing.T) {
	dsn := "jbc:123456@tcp(47.96.86.248:3306)/mydb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(db)

}

func TestDb2(t *testing.T) {
	b := userDao.CheckUser(123, 0, "2212321")
	fmt.Println(b)
}

func TestAddUser(t *testing.T) {
	user := &database.User{
		Id:         util.GetSnowflakeId(),
		Username:   "lyz",
		Password:   "03",
		Phone:      138,
		Type:       0,
		CreateTime: time.Now(),
		IsDelete:   0,
		Ext:        sql.NullString{},
	}
	b := userDao.AddUser(user)
	fmt.Println(b)
}

func TestAddImage(t *testing.T) {
	filebytes, err := ioutil.ReadFile("C:\\Users\\78240\\go\\ticket\\image\\2e3631ccbfd2850a668ece07ff63811d.jpeg")
	if err != nil {
		fmt.Println(err)
	}
	img := &database.Image{
		Id:         util.GetSnowflakeId(),
		TicketId:   util.GetSnowflakeId(),
		BinaryData: filebytes,
		Type:       0,
		CreateTime: time.Now(),
	}
	b := imageDao.AddImage(img)
	fmt.Println(b)
}

func TransToBase64() {
	//base64.StdEncoding.DecodeString(datasource)
	//emptyBuff := bytes.NewBuffer(nil) //开辟一个新的空buff
	//filebytes, err := ioutil.ReadFile(dst)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//jpeg.Encode(emptyBuff, subImg, nil)                //img写入到buff
	//dist := make([]byte, 50000)                        //开辟存储空间
	//base64.StdEncoding.Encode(dist, emptyBuff.Bytes()) //buff转成base64
	//fmt.Println(string(dist))                          //输出图片base64(type = []byte)
}

package test

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
	"ticket/database"
	"ticket/util"
	"time"
)

var userDao database.UserImpl

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

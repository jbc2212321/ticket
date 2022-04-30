package database

import "fmt"

type Verifylist struct {
	ListId   int    `gorm:"column:list_id" db:"list_id" json:"list_id" form:"list_id"`     //审核列表 id
	Songid   int    `gorm:"column:songid" db:"songid" json:"songid" form:"songid"`         //歌曲id
	Songname string `gorm:"column:songname" db:"songname" json:"songname" form:"songname"` //歌曲名
	Userid   int64  `gorm:"column:userid" db:"userid" json:"userid" form:"userid"`         //用户id
	Username string `gorm:"column:username" db:"username" json:"username" form:"username"` //用户名
	Status   int    `gorm:"column:status" db:"status" json:"status" form:"status"`         //申请状态
}

func (m *Verifylist) TableName() string {
	return "verifylist"
}

type VerifylistImpl struct {
}

//func (u *UserImpl) AddUser(user *User) error {
//	db := GetDB()
//	err := db.Create(&user).Error
//	if err != nil {
//		return err
//	}
//	return nil
//}

func (*VerifylistImpl) GetVerifyByID(idList ...int64) ([]*Verifylist,error){
	db := GetDB()
	veris:=new([]*Verifylist)
	err := db.Model(&Verifylist{}).Where("userid = ?", idList).Find(veris).Error
	if err != nil {
		return nil, err
	}
	return *veris,nil
}

func (*VerifylistImpl) GetVerifyByStatus(statuList ...int64) ([]*Verifylist,error){
	db := GetDB()
	veris:=new([]*Verifylist)
	fmt.Println("statuList:",statuList)
	err := db.Model(&Verifylist{}).Where("status = ?", statuList).Find(veris).Error
	if err != nil {
		return nil, err
	}
	return *veris,nil
}
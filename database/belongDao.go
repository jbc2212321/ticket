package database

type Belong struct {
	Num      int    `gorm:"column:num" db:"num" json:"num" form:"num"`                     //编号
	Userid   int64  `gorm:"column:userid" db:"userid" json:"userid" form:"userid"`         //用户id
	Songid   int    `gorm:"column:songid" db:"songid" json:"songid" form:"songid"`         //歌曲id
	Username string `gorm:"column:username" db:"username" json:"username" form:"username"` //用户名
	Songname string `gorm:"column:songname" db:"songname" json:"songname" form:"songname"` //歌曲名
}

func (m *Belong) TableName() string {
	return "belong"
}

type BelongImpl struct {
}

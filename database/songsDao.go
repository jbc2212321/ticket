package database

type Songs struct {
	Id     int    `gorm:"column:id" db:"id" json:"id" form:"id"`         //歌曲ID
	Name   string `gorm:"column:name" db:"name" json:"name" form:"name"` //歌曲名称
	Status int    `gorm:"column:status" db:"status" json:"status" form:"status"`
}

func (m *Songs) TableName() string {
	return "songs"
}

type SongsImpl struct {
}

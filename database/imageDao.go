package database

import "time"

type Image struct {
	Id         int64     `gorm:"column:id" db:"id" json:"id" form:"id"`                                     //图片id
	TicketId   int64     `gorm:"column:ticket_id" db:"ticket_id" json:"ticket_id" form:"ticket_id"`         //小票id
	Type       int       `gorm:"column:type" db:"type" json:"type" form:"type"`                             //小票类型
	BinaryData []byte    `gorm:"column:binary_data" db:"binary_data" json:"binary_data" form:"binary_data"` //图片2进制文件
	CreateTime time.Time `gorm:"column:create_time" db:"create_time" json:"create_time" form:"create_time"` //创建日期
}

package database

import "time"

type Train struct {
	Id                   int64     `gorm:"column:id" db:"id" json:"id" form:"id"`                                                                             //发票id
	UserId               int64     `gorm:"column:user_id" db:"user_id" json:"user_id" form:"user_id"`                                                         //提交人id
	CreateTime           time.Time `gorm:"column:create_time" db:"create_time" json:"create_time" form:"create_time"`                                         //创建时间
	IdentificationNumber string    `gorm:"column:Identification_number" db:"Identification_number" json:"Identification_number" form:"Identification_number"` //识别号码
	TicketNum            string    `gorm:"column:ticket_num" db:"ticket_num" json:"ticket_num" form:"ticket_num"`                                             //车票号
	StartingStation      string    `gorm:"column:starting_station" db:"starting_station" json:"starting_station" form:"starting_station"`                     //始发站
	TrainNum             string    `gorm:"column:train_num" db:"train_num" json:"train_num" form:"train_num"`                                                 //车次号
	DestinationStation   string    `gorm:"column:destination_station" db:"destination_station" json:"destination_station" form:"destination_station"`         //到达站
	Date                 string    `gorm:"column:date" db:"date" json:"date" form:"date"`                                                                     //到达日期
	TicketRates          string    `gorm:"column:ticket_rates" db:"ticket_rates" json:"ticket_rates" form:"ticket_rates"`                                     //车票金额
	SeatCategory         string    `gorm:"column:seat_category" db:"seat_category" json:"seat_category" form:"seat_category"`                                 //席别
	Name                 string    `gorm:"column:name" db:"name" json:"name" form:"name"`                                                                     //姓名
	IdNum                string    `gorm:"column:id_num" db:"id_num" json:"id_num" form:"id_num"`                                                             //身份证号码
	SerialNumber         string    `gorm:"column:serial_number" db:"serial_number" json:"serial_number" form:"serial_number"`                                 //序列号
	SalesStation         string    `gorm:"column:sales_station" db:"sales_station" json:"sales_station" form:"sales_station"`                                 //售站
	Time                 string    `gorm:"column:time" db:"time" json:"time" form:"time"`                                                                     //时间
	SeatNum              string    `gorm:"column:seat_num" db:"seat_num" json:"seat_num" form:"seat_num"`                                                     //座位号
}

func (t *Train) TableName() string {
	return "train"
}

type trainImpl struct {
}

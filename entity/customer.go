package entity

import "time"

type Customer struct {
	Row      int32     `gorm:"column:row;primaryKey"`
	Id       string    `gorm:"column:id"`
	Name     string    `gorm:"column:name"`
	Birthday time.Time `gorm:"column:birthday"`
	Age      int32     `gorm:"column:age"`
	Remark   string    `gorm:"column:remark"`
}

func (i *Customer) TableName() string {
	return "customer"
}

package entity

type Address struct {
	Row  int32  `gorm:"column:row;primaryKey"`
	Id   string `gorm:"column:id"`
	Text string `gorm:"column:text"`
}

func (i *Address) TableName() string {
	return "address"
}

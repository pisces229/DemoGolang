package dto

type DefaultPersonDto struct {
	A int32  `gorm:"column:a"`
	B string `gorm:"column:b"`
	C string `gorm:"column:c"`
}

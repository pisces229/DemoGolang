package entity

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Person struct {
	Row      int32     `gorm:"column:row;primaryKey"`
	Id       string    `gorm:"column:id"`
	Name     string    `gorm:"column:name"`
	Age      int32     `gorm:"column:age"`
	Birthday time.Time `gorm:"column:birthday"`
	Remark   string    `gorm:"column:remark"`
}

func (i *Person) TableName() string {
	return "person"
}

func (i *Person) BeforeSave(db *gorm.DB) (err error) {
	fmt.Println("BeforeSave")
	// err = errors.New("do BeforeSave error")
	return
}

func (i *Person) BeforeCreate(db *gorm.DB) (err error) {
	fmt.Println("BeforeCreate")
	// err = errors.New("do BeforeCreate error")
	return
}

func (i *Person) AfterCreate(db *gorm.DB) (err error) {
	fmt.Println("AfterCreate")
	// err = errors.New("do AfterCreate error")
	return
}

func (i *Person) AfterSave(db *gorm.DB) (err error) {
	fmt.Println("AfterSave")
	// err = errors.New("do AfterSave error")
	return
}

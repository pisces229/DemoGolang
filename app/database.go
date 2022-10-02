package app

import (
	"context"
	"demo.golang/entity"
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func databaseSampleQuery(db *gorm.DB) {
	{
		data := &entity.Person{}
		db.Model(entity.Person{}).First(data)
		fmt.Printf("%+v \n", data)
	}
	{
		data := &[]entity.Person{}
		// data := &[]map[string]interface{}{}
		db.Model(entity.Person{}).Find(data, 1)
		//db.Model(entity.Address{}).Find(data, []int{1, 2, 3})
		fmt.Printf("%+v \n", data)
	}
	{
		data := &[]entity.Person{}
		db.Model(entity.Person{}).
			// Select().
			Where("row > ?", 0).
			// Where("row > ?", 0).
			// Or("row > ?", 0).
			Order("row").
			// Order("row ASC").
			// Order("row DESC").
			Limit(2).
			Find(data)
		fmt.Printf("%+v \n", data)
	}
	type databaseScanDto struct {
		A int32  `gorm:"column:a"`
		B string `gorm:"column:b"`
		C string `gorm:"column:c"`
	}
	{
		type param struct {
			Row int32
		}
		dto := &[]databaseScanDto{}
		db.
			// Raw("SELECT row as a, id as b, text as c FROM address WHERE row > ?", 0).
			// Raw("SELECT row as a, id as b, text as c FROM address WHERE row > @Row", sql.Named("Row", 0)).
			// Raw("SELECT row as a, id as b, text as c FROM address WHERE row > @Row", map[string]interface{}{"Row": 0}).
			Raw("SELECT row as a, id as b, text as c FROM address WHERE row > @Row", param{Row: 0}).
			Scan(dto)
		fmt.Printf("%+v \n", dto)
	}
	{
		rows, _ := db.
			Raw("SELECT row as a, id as b, text as c FROM address WHERE row > ?", 0).
			Rows()
		defer func() {
			if err := rows.Close(); err != nil {
				fmt.Println(err)
			}
		}()
		var a int32
		var b string
		var c string
		dto := databaseScanDto{}
		for rows.Next() {
			if err := rows.Scan(&a, &b, &c); err == nil {
				fmt.Printf("[%+v][%+v][%+v] \n", a, b, c)
				if err := db.ScanRows(rows, &dto); err == nil {
					fmt.Printf("%+v \n", dto)
				} else {
					fmt.Println(err)
				}
			} else {
				fmt.Println(err)
			}
		}
	}
}
func databaseSampleCreate(db *gorm.DB) {
	data := entity.Person{
		Id:       uuid.New().String(),
		Name:     "Create",
		Age:      18,
		Birthday: time.Now(),
		Remark:   uuid.New().String(),
	}
	db.Create(&data)
	fmt.Printf("%+v \n", data)
}
func databaseSampleModify(db *gorm.DB) {
	data := entity.Person{}
	db.Model(entity.Person{}).Where("id = ?", 0).First(&data)
	data.Remark = uuid.New().String()
	db.Save(&data)
	//db.Model(&entity.Person{}).Where("id = ?", 0).
	//	Update("Remark", uuid.New().String())
	//db.Model(&entity.Person{}).Where("id = ?", 0).
	//	Updates(&entity.Person{Name: "Modify", Remark: uuid.New().String()})
	//db.Model(&entity.Person{}).Where("id = ?", 0).
	//	Updates(map[string]interface{}{"name": "Modify", "remark": uuid.New().String()})
	fmt.Printf("%+v \n", data)
}
func databaseSampleRemove(db *gorm.DB) {
	data := entity.Person{}
	db.Model(entity.Person{}).Where("id = ?", 0).First(&data)
	db.Delete(&data)
	//db.Where("id = ?", 0).Delete(entity.Person{})
	fmt.Printf("%+v \n", data)
}
func databaseSampleTransaction(db *gorm.DB) {
	err := db.Transaction(func(tx *gorm.DB) error {
		tx.Model(entity.Person{}).First(&entity.Person{})
		// if err := tx.Create(&Person{Id: "0", Name: "Trans", Age: 18, Birthday: time.Now(), Remark: uuid.New().String()}).Error; err != nil {
		// 	return err
		// }
		// data := Person{}
		// tx.Model(Person{}).Where("id = '0'").First(&data)
		// fmt.Printf("%+v \n", data)
		// if err := tx.Create(&Person{Id: "0", Name: "Trans", Age: 18, Birthday: time.Now(), Remark: uuid.New().String()}).Error; err != nil {
		// 	return err
		// }
		return nil
	})
	if err != nil {
		fmt.Println(err.Error())
	}
}
func databaseSamplePool(db *gorm.DB) {
	// f := func(i int) {
	// 	_, err := db.Raw("select 1").Rows()
	// 	time.Sleep(1 * time.Second)
	// 	if err != nil {
	// 		fmt.Println(i, err.Error())
	// 	} else {
	// 		fmt.Println(i, "success")
	// 	}
	// }
	// f := func(i int) {
	// 	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	// 	defer cancel()
	// 	_, err := db.WithContext(ctx).Raw("select 1").Rows()
	// 	time.Sleep(1 * time.Second)
	// 	if err != nil {
	// 		fmt.Println(i, err.Error())
	// 	} else {
	// 		fmt.Println(i, "success")
	// 	}
	// }
	f := func(i int) {
		// ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		// defer cancel()
		// err := db.WithContext(ctx).Transaction(func(tx *gorm.DB) error { ...
		err := db.Transaction(func(tx *gorm.DB) error {
			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancel()
			if _, err := tx.WithContext(ctx).Raw("select 1").Rows(); err != nil {
				fmt.Println(i, err.Error())
			} else {
				time.Sleep(1 * time.Second)
				fmt.Println(i, "success")
			}
			// tx.Model(Person{}).First(&Person{})
			// time.Sleep(1 * time.Second)
			return nil
		})
		if err != nil {
			fmt.Println(i, err.Error())
		}
	}
	for i := 0; i < 10; i++ {
		go f(i)
	}
	time.Sleep(1 * time.Second)
}
func databaseSampleSession(db *gorm.DB) {
	sessionLogger := logrus.New()
	sessionLogger.SetOutput(os.Stdout)
	sessionLogger.SetFormatter(&logrus.JSONFormatter{})
	sessionLogger.SetLevel(logrus.InfoLevel)
	tx := db.Session(&gorm.Session{Logger: logger.New(
		log.New(sessionLogger.WithFields(logrus.Fields{"uuid": uuid.New()}).Writer(), "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)})
	fmt.Println(tx)
}

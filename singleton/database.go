package singleton

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func initDatabase() *gorm.DB {
	var dns string
	var dialector gorm.Dialector
	// Sqlserver
	// dns = fmt.Sprintf(
	// 	"Server=%s;Database=%s;User ID=%s;Password=%s;",
	// 	SingletonConfiguration.Connection.Sqlserver.Server,
	// 	SingletonConfiguration.Connection.Sqlserver.Database,
	// 	SingletonConfiguration.Connection.Sqlserver.Username,
	// 	SingletonConfiguration.Connection.Sqlserver.Password)
	// dialector = sqlserver.New(sqlserver.Config{
	// 	DSN: dns,
	// })
	// Postgres
	// dns = fmt.Sprintf(
	// 	"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
	// 	SingletonConfiguration.Connection.Postgres.Host,
	// 	SingletonConfiguration.Connection.Postgres.Username,
	// 	SingletonConfiguration.Connection.Postgres.Password,
	// 	SingletonConfiguration.Connection.Postgres.Database,
	// 	SingletonConfiguration.Connection.Postgres.Port)
	//dialector = postgres.New(postgres.Config{
	//	DSN: dns,
	//	// PreferSimpleProtocol: true,
	//})
	fmt.Println(dns)
	dbLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		// log.New(databaseLogFile, "\r\n", log.LstdFlags),
		// log.New(AppLogrus.Writer(), "\r\n", log.LstdFlags),
		// AppLogwriter,
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		})
	db, err := gorm.Open(dialector,
		&gorm.Config{
			// SkipDefaultTransaction: true,
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
				NoLowerCase:   true,
			},
			Logger: dbLogger,
			// DryRun: false,
			// PrepareStmt: false,
			// DisableAutomaticPing: true,
			// DisableForeignKeyConstraintWhenMigrating: true,
			CreateBatchSize: 1000,
		})
	if err != nil {
		panic("failed to connect database")
	}
	if db, err := db.DB(); err != nil {
		panic("get sql db failed:" + err.Error())
	} else {
		db.SetConnMaxLifetime(time.Duration(SingletonConfiguration.Connection.Pool.MaxLifetime) * time.Second)
		db.SetMaxIdleConns(SingletonConfiguration.Connection.Pool.MaxIdleConns)
		db.SetMaxOpenConns(SingletonConfiguration.Connection.Pool.MaxOpenConns)
	}
	return db
}

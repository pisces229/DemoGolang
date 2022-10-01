package singleton

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var AppDatabase = initDatabase()

func initDatabase() (db *gorm.DB) {
	fmt.Println("initDatabase")
	// db = initDatabaseSqlserver()
	db = initDatabasePostgres()
	return
}

func initDatabaseSqlserver() (db *gorm.DB) {
	dns := fmt.Sprintf("Server=%s;Database=%s;User ID=%s;Password=%s;",
		AppConfiguration.Connection.Sqlserver.Server,
		AppConfiguration.Connection.Sqlserver.Database,
		AppConfiguration.Connection.Sqlserver.Username,
		AppConfiguration.Connection.Sqlserver.Password)
	fmt.Println(dns)
	// var databaseLogFile, _ = os.Create("database.log")
	db, err := gorm.Open(sqlserver.Open(dns),
		&gorm.Config{
			// SkipDefaultTransaction: true,
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
				NoLowerCase:   true,
			},
			Logger: logger.New(
				log.New(os.Stdout, "\r\n", log.LstdFlags),
				// log.New(databaseLogFile, "\r\n", log.LstdFlags),
				// log.New(AppLogrus.Writer(), "\r\n", log.LstdFlags),
				// AppLogwriter,
				logger.Config{
					SlowThreshold:             time.Second,
					LogLevel:                  logger.Info,
					IgnoreRecordNotFoundError: true,
					Colorful:                  true,
				},
			),
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
		db.SetConnMaxLifetime(time.Duration(AppConfiguration.Connection.Pool.MaxLifetime) * time.Second)
		db.SetMaxIdleConns(AppConfiguration.Connection.Pool.MaxIdleConns)
		db.SetMaxOpenConns(AppConfiguration.Connection.Pool.MaxOpenConns)
	}
	return
}

func initDatabasePostgres() (db *gorm.DB) {
	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		AppConfiguration.Connection.Postgres.Host,
		AppConfiguration.Connection.Postgres.Username,
		AppConfiguration.Connection.Postgres.Password,
		AppConfiguration.Connection.Postgres.Database,
		AppConfiguration.Connection.Postgres.Port)
	fmt.Println(dns)
	// var databaseLogFile, _ = os.Create("database.log")
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dns,
		// PreferSimpleProtocol: true,
	}), &gorm.Config{
		// SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			// TablePrefix: "t_",
			SingularTable: true,
			NoLowerCase:   true,
		},
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			// log.New(databaseLogFile, "\r\n", log.LstdFlags),
			// log.New(AppLogrus.Writer(), "\r\n", log.LstdFlags),
			// AppLogwriter,
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  logger.Info,
				IgnoreRecordNotFoundError: true,
				Colorful:                  true,
			},
		),
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
		db.SetConnMaxLifetime(time.Duration(AppConfiguration.Connection.Pool.MaxLifetime) * time.Second)
		db.SetMaxIdleConns(AppConfiguration.Connection.Pool.MaxIdleConns)
		db.SetMaxOpenConns(AppConfiguration.Connection.Pool.MaxOpenConns)
	}
	return
}

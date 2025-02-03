package db

import (
	"DB-SETUP/config"
	"fmt"
	"log"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB   *gorm.DB
	once sync.Once
)

func InitDB() {
	once.Do(func() {
		cfg, err := config.LoadConfig()
		if err != nil {
			log.Fatalf("error in set up %v", err)
		}

		dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=%s",
			cfg.Database.Host,
			cfg.Database.Port,
			cfg.Database.User,
			cfg.Database.Password,
			cfg.Database.SSLMode)

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
		if err != nil {
			log.Fatalf("error in connecting to db: %v", err)
		}

		createDBSQL := fmt.Sprintf("CREATE DATABASE %s;", cfg.Database.DBName)
		db.Exec(createDBSQL)

		dsn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			cfg.Database.Host,
			cfg.Database.Port,
			cfg.Database.User,
			cfg.Database.Password,
			cfg.Database.DBName,
			cfg.Database.SSLMode)

		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("error in connecting new db: %v", err)
		}

		DB = db
		fmt.Println("successfully connected to database")
	})
}

func GetDB() *gorm.DB {
	if DB == nil {
		InitDB()
	}
	return DB
}

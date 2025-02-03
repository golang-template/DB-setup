package db

import (
	"DB-SETUP/models"
	"fmt"
)

func Migrate() {
	db := GetDB()
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		fmt.Println("error in migrating", err)
	} else {
		fmt.Println("migrations completed")
	}
}

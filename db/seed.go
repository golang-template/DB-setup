package db

import "DB-SETUP/models"

func Seed() {
	db := GetDB()

	db.Create(&models.User{FirstName: "admin", LastName: "admin"})
}

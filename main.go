package main

import (
	"DB-SETUP/db"
	"DB-SETUP/models"
	"fmt"
)

func main() {
	db.InitDB()
	db.Migrate()

	var firstName, lastName string
	fmt.Print("plz enter your first name: ")
	fmt.Scanln(&firstName)
	fmt.Print("plz enter your last name: ")
	fmt.Scanln(&lastName)

	user := models.User{FirstName: firstName, LastName: lastName}
	database := db.GetDB()
	result := database.Create(&user)

	if result.Error != nil {
		fmt.Println("error in saving", result.Error)
	} else {
		fmt.Println("user saved successfully", user.FirstName, user.LastName)
	}
}

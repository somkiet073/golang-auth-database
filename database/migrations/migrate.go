package migrations

import (
	"fmt"
	"golang-auth-database/database/db"
	"golang-auth-database/database/models"
)

func MigrateDB() {
	db := db.Connects()
	type User models.User
	type Department models.Department
	type Positions models.Positions

	db.AutoMigrate(&User{}, &Department{}, &Positions{})

	fmt.Println("Create Database Success.")
	defer db.Close()
}

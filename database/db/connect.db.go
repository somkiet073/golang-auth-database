package db

import (
	"golang-auth-database/helpers"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// _ "github.com/go-sql-driver/mysql"
)

func Connects() *gorm.DB {
	dbDriver := "mysql"
	dbHost := "127.0.0.1"
	dbPort := "3306"
	dbName := "godb"
	dbUser := "root"
	dbPass := ""
	// "github.com/go-sql-driver/mysql"
	// db, err = sql.Open("mysql", "<root>:<password>@tcp(localhost:5555)/<dbname>")
	db, err := gorm.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbHost+":"+dbPort+")/"+dbName)
	helpers.HandleErr(err)
	return db
}

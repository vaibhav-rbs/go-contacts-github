package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB //database

func init() {
	var err error
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	// defer db.Close()
	db.Debug().AutoMigrate(&Account{}, &Contact{}) //Database migration
}

//returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}

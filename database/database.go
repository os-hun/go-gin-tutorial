package database

import (
	"github.com/jinzhu/gorm"
	"go-gin-tutorial/model"
	_"github.com/mattn/go-sqlite3"
)

func DbInit() {
	// open db
	db, err := gorm.Open("sqlite3", "test.splite3")
	if err != nil {
		panic("can't open the database")
	}
	
	// auto migration
	db.AutoMigrate(&model.Todo{})
	
	// close db
	defer db.Close()
}

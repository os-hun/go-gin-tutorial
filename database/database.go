package database

import (
	"github.com/jinzhu/gorm"
	_"github.com/mattn/go-sqlite3"
)

func DbInit() {
	db, err := gorm.Open("sqlite3", "test.splite3")
	if err != nil {
		panic("can't open the database")
	}
	defer db.Close()
}

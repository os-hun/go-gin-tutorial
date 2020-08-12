package database

import (
	"github.com/jinzhu/gorm"
	"go-gin-tutorial/model"
	_"github.com/mattn/go-sqlite3"
)

func DbInit() {
	// open db
	db, err := gorm.Open("sqlite3", "database/test.splite3")
	if err != nil {
		panic("データベースへの接続に失敗しました")
	}

	// auto migration
	db.AutoMigrate(&model.Todo{})

	// close db
	defer db.Close()
}

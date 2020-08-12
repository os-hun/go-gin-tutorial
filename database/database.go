package database

import (
	"github.com/jinzhu/gorm"
	"go-gin-tutorial/model"
	_"github.com/mattn/go-sqlite3"
)

// 初期化
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

// 追加
func DbInsert(text, status string) {
	db, err := gorm.Open("sqlite3", "database/test.splite3")
	if err != nil {
		panic("データベースへの接続に失敗しました。")
	}
	db.Create(&model.Todo{Text: text, Status: status})
	defer db.Close()
}

// 更新
func DbUpdate(id int, text, status string) {
	db, err := gorm.Open("sqlite3", "database/test.splite3")
	if err != nil {
		panic("データベースへの接続に失敗しました。")
	}
	var todo model.Todo
	db.First(&todo, id)
	todo.Text = text
	todo.Status = status
	db.Save(&todo)
	db.Close()
}

// 削除
func DbDelete(id int) {
	db, err := gorm.Open("sqlite3", "database/test.splite3")
	if err != nil {
		panic("データベースへの接続に失敗しました。")
	}
	var todo model.Todo
	db.First(&todo, id)
	db.Delete(&todo)
	db.Close()
}

// 全取得
func DbGetAll() []model.Todo {
	db, err := gorm.Open("sqlite3", "database/test.splite3")
	if err != nil {
		panic("データベースへの接続に失敗しました。")
	}
	var todos []model.Todo
	db.Order("created_at desc").Find(&todos)
	db.Close()
	return todos
}

// １つ取得
func DbGetOne(id int) model.Todo {
	db, err := gorm.Open("sqlite3", "database/test.splite3")
	if err != nil {
		panic("データベースへの接続に失敗しました。")
	}
	var todo model.Todo
	db.First(&todo, id)
	db.Close()
	return todo
}

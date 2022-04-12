package model

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

type Todo struct {
	// 標準のモデルを呼び出し
	gorm.Model
	Text   string
	Status string
}

// DB接続
func sqlConnect() (database *gorm.DB) {
	DBMS := "mysql"
	USER := "root"
	PASS := "root"
	PROTOCOL := "tcp(db:3306)"
	DBNAME := "go_database"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"

	count := 0
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		for {
			if err == nil {
				fmt.Println("")
				break
			}
			fmt.Print(".")
			time.Sleep(time.Second)
			count++
			if count > 180 {
				fmt.Println("")
				fmt.Println("DB接続失敗")
				panic(err)
			}
			db, err = gorm.Open(DBMS, CONNECT)
		}
	}
	fmt.Println("DB接続成功")

	return db
}

// DB初期化
func Init() {
	db := sqlConnect()
	db.AutoMigrate(&Todo{}) // マイグレートを実行
	defer db.Close()
}

// DB追加
func Insert(text, status string) {
	db := sqlConnect()
	db.Create(&Todo{Text: text, Status: status}) // 引数追加
	defer db.Close()
}

// DB全取得
func SelectAll() []Todo {
	db := sqlConnect()
	var todos []Todo
	db.Order("created_at desc").Find(&todos) // Find(&todos)：構造体Todoに対するデーブルの要素全てを取得、Order("created_at desc")：新しい順
	db.Close()
	return todos
}

// DB一つ取得
func SelectOne(id int) Todo {
	db := sqlConnect()
	var todo Todo
	db.First(&todo, id)
	db.Close()
	return todo
}

// DB更新
func Update(id int, text, status string) {
	db := sqlConnect()
	var todo Todo
	db.First(&todo, id)
	todo.Text = text
	todo.Status = status
	db.Save(&todo)
	db.Close()
}

// DB削除
func Delete(id int) {
	db := sqlConnect()
	var todo Todo
	db.First(&todo, id)
	db.Delete(&todo)
	db.Close()
}

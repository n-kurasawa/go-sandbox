package main

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	Name      string
	Age       uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

func main() {
	db, err := gorm.Open("mysql", "root:root@(localhost)/gorm_sample?charset=utf8&parseTime=True")
	if err != nil {
		panic("データベースへの接続に失敗しました")
	}
	defer db.Close()
	db.LogMode(true)

	utc, _ := time.LoadLocation("UTC")
	jst, _ := time.LoadLocation("Asia/Tokyo")

	createdAt := time.Date(2022, 7, 1, 0, 0, 0, 0, utc)
	updatedAt := createdAt.In(jst)

	// Create
	db.Create(&User{Name: "test", Age: 20, CreatedAt: createdAt, UpdatedAt: updatedAt})

	// Read
	user := User{}
	db.First(&user, "name = ?", "test")

	fmt.Println(user)
}

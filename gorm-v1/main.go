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

	jstDate := time.Date(2022, 7, 1, 0, 0, 0, 0, jst)
	utcDate := jstDate.In(utc)

	// Create
	db.Create(&User{Name: "test", Age: 20, CreatedAt: jstDate, UpdatedAt: utcDate})

	// Read
	type CreatedAt struct {
		CreatedAt time.Time
	}
	var jstCreatedAt CreatedAt
	jstSQL := `
		select created_at from users where date(convert_tz( created_at, '+00:00', 'Asia/Tokyo' )) = ?
	`
	if err := db.Raw(jstSQL, jstDate).Scan(&jstCreatedAt).Error; err != nil {
		fmt.Println("jst err: ", err)
	} else {
		fmt.Println("jst: ", jstCreatedAt)
	}

	var utcCreatedAt CreatedAt
	utcSQL := `
		select created_at from users where created_at = ?
	`
	if err := db.Raw(utcSQL, utcDate).Scan(&utcCreatedAt).Error; err != nil {
		fmt.Println("utc err: ", err)
	} else {
		fmt.Println("utc: ", utcCreatedAt.CreatedAt)
	}
}

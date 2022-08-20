package main

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type sample struct {
	ID       int
	Datetime time.Time
}

func (sample) TableName() string {
	return "sample_table"
}

func main() {
	db, err := gorm.Open(mysql.Open("root:root@/sample?parseTime=true"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("unable to use data source name", err)
	}

	utc, _ := time.LoadLocation("UTC")
	utcTime := time.Date(2022, 8, 20, 0, 0, 0, 0, utc)
	var utcResult sample
	db.Where("datetime = ?", utcTime).First(&utcResult)
	fmt.Printf("[1] id: %d, datetime: %v\n", utcResult.ID, utcResult.Datetime)

	jst, _ := time.LoadLocation("Asia/Tokyo")
	jstTime := time.Date(2022, 8, 20, 0, 0, 0, 0, jst)
	var jstResult sample
	if err := db.Where("datetime = ?", jstTime).First(&jstResult).Error; err == nil {
		fmt.Printf("[1] id: %d, datetime: %v\n", jstResult.ID, jstResult.Datetime)
	} else {
		fmt.Println(err)
	}
}

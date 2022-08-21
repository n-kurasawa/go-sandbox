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

	// [1] UTC の 2022-08-20 00:00:00 で検索
	utc, _ := time.LoadLocation("UTC")
	utcTime := time.Date(2022, 8, 20, 0, 0, 0, 0, utc)
	if result, err := findByDatetime(db, utcTime); err == nil {
		fmt.Printf("[1] id: %d, datetime: %v\n", result.ID, result.Datetime)
	} else {
		fmt.Println("[1]", err)
	}

	// [2] JST の 2022-08-20 00:00:00 で検索
	jst, _ := time.LoadLocation("Asia/Tokyo")
	jstTime := time.Date(2022, 8, 20, 0, 0, 0, 0, jst)
	if result, err := findByDatetime(db, jstTime); err == nil {
		fmt.Printf("[2] id: %d, datetime: %v\n", result.ID, result.Datetime)
	} else {
		fmt.Println("[2]", err)
	}

	// [3] UTC の 2022-08-20 00:00:00 を JST にして検索
	if result, err := findByDatetime(db, utcTime.In(jst)); err == nil {
		fmt.Printf("[3] id: %d, datetime: %v\n", result.ID, result.Datetime)
	} else {
		fmt.Println("[3]", err)
	}
}

func findByDatetime(db *gorm.DB, time time.Time) (sample, error) {
	var result sample
	if err := db.Where("datetime = ?", time).First(&result).Error; err != nil {
		return sample{}, nil
	}
	return result, nil
}

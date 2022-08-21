package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type sample struct {
	id       int
	datetime time.Time
}

func main() {
	db, err := sql.Open("mysql", "root:root@/sample?parseTime=true")
	if err != nil {
		log.Fatal("unable to use data source name", err)
	}
	defer db.Close()

	// [1] UTC の 2022-08-20 00:00:00 で検索
	utc, _ := time.LoadLocation("UTC")
	utcTime := time.Date(2022, 8, 20, 0, 0, 0, 0, utc)
	if result, err := findByDatetime(db, utcTime); err == nil {
		fmt.Printf("[1] id: %d, datetime: %v\n", result.id, result.datetime)
	} else {
		fmt.Println("[1]", err)
	}

	// [2] JST の 2022-08-20 00:00:00 で検索
	jst, _ := time.LoadLocation("Asia/Tokyo")
	jstTime := time.Date(2022, 8, 20, 0, 0, 0, 0, jst)
	if result, err := findByDatetime(db, jstTime); err == nil {
		fmt.Printf("[2] id: %d, datetime: %v\n", result.id, result.datetime)
	} else {
		fmt.Println("[2]", err)
	}

	// [3] UTC の 2022-08-20 00:00:00 を JST にして検索
	if result, err := findByDatetime(db, utcTime.In(jst)); err == nil {
		fmt.Printf("[3] id: %d, datetime: %v\n", result.id, result.datetime)
	} else {
		fmt.Println("[3]", err)
	}
}

func findByDatetime(db *sql.DB, datetime time.Time) (sample, error) {
	var result sample
	if err := db.QueryRow("select id, datetime from sample_table where datetime = ?", datetime).Scan(&result.id, &result.datetime); err != nil {
		return sample{}, err
	}
	return result, nil
}

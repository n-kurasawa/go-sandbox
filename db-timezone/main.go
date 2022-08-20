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

	utc, _ := time.LoadLocation("UTC")
	utcDatetime := time.Date(2022, 8, 20, 0, 0, 0, 0, utc)
	if result, err := findByDatetime(db, utcDatetime); err == nil {
		fmt.Printf("[1] id: %d, datetime: %v\n", result.id, result.datetime)
	} else {
		fmt.Println("[1]", err)
	}

	jst, _ := time.LoadLocation("Asia/Tokyo")
	jstDatetime := time.Date(2022, 8, 20, 0, 0, 0, 0, jst)
	if result, err := findByDatetime(db, jstDatetime); err == nil {
		fmt.Printf("[2] id: %d, datetime: %v\n", result.id, result.datetime)
	} else {
		fmt.Println("[2]", err)
	}

	if result, err := findByDatetime(db, utcDatetime.In(jst)); err == nil {
		fmt.Printf("[3] id: %d, datetime: %v\n", result.id, result.datetime)
	} else {
		fmt.Println("[3]", err)
	}
}

func findByDatetime(db *sql.DB, dateTime time.Time) (sample, error) {
	var result sample
	if err := db.QueryRow("select id, datetime from sample_table where datetime = ?", dateTime).Scan(&result.id, &result.datetime); err != nil {
		return sample{}, fmt.Errorf("failed QueryRow: %w", err)
	}
	return result, nil
}

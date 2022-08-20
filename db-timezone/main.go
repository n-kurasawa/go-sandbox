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

	var utcResult sample
	if err := db.QueryRow("select id, datetime from sample_table where datetime = ?", utcDatetime).Scan(&utcResult.id, &utcResult.datetime); err != nil {
		log.Fatal("failed QueryRow: ", err)
	}
	fmt.Printf("id: %d, datetime: %v\n", utcResult.id, utcResult.datetime)

	jst, _ := time.LoadLocation("Asia/Tokyo")
	jstDatetime := time.Date(2022, 8, 20, 9, 0, 0, 0, jst)
	var jstResult sample
	if err := db.QueryRow("select id, datetime from sample_table where convert_tz(datetime, 'UTC', 'Asia/Tokyo') = ?", jstDatetime).Scan(&jstResult.id, &jstResult.datetime); err != nil {
		log.Fatal("failed QueryRow: ", err)
	}
	fmt.Printf("id: %d, datetime: %v\n", jstResult.id, jstResult.datetime)
}

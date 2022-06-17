package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	db, err := sql.Open("pgx", "host=localhost port=5432 user=testuser dbname=testdb password=pass sslmode=disable")
	if nil != err {
		log.Fatal(err)
	}

	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err)
	}

	s := &Service{db: db}
	user := User{UserID: "2", UserName: "test2", CreatedAt: time.Now()}
	if err := s.CreateUser(ctx, user); err != nil {
		fmt.Println(err.Error())
	}
}

type User struct {
	UserID    string
	UserName  string
	CreatedAt time.Time
}

type Service struct {
	db *sql.DB
}

func (s *Service) CreateUser(ctx context.Context, user User) (err error) {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	if _, err = tx.ExecContext(ctx, `INSERT INTO users (user_id, user_name, created_at) VALUES ($1, $2, $3)`, user.UserID, user.UserName, user.CreatedAt); err != nil {
		return err
	}

	return tx.Commit()
}

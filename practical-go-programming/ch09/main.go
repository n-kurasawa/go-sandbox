package main

import (
	"context"
	"database/sql"
	"errors"
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

	s := &Service{tx: txAdmin{db}}
	if err := s.UpdateProduct(ctx); err != nil {
		log.Fatal(err)
	}
}

type User struct {
	UserID    string
	UserName  string
	CreatedAt time.Time
}

// transaction-wrapper-start
// txAdminはトランザクション制御するための構造体
type txAdmin struct {
	*sql.DB
}

type Service struct {
	tx txAdmin
}

// Transaction はトランザクションを制御するメソッド
// アプリケーション開発者が本メソッドを使って、DMLのクエリーを発行する
func (t *txAdmin) Transaction(ctx context.Context, f func(ctx context.Context, tx *sql.Tx) (err error)) error {
	tx, err := t.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	if err := f(ctx, tx); err != nil {
		return fmt.Errorf("transaction query failed: %w", err)
	}
	return tx.Commit()
}

func (s *Service) UpdateProduct(ctx context.Context) error {
	updateFunc := func(ctx context.Context, tx *sql.Tx) error {
		if _, err := s.tx.ExecContext(ctx, `UPDATE users SET user_name = 'upd123' WHERE user_id = '1';`); err != nil {
			return err
		}
		if _, err := s.tx.ExecContext(ctx, `UPDATE users SET user_name = 'upd234' WHERE user_id = '2';`); err != nil {
			return err
		}
		return errors.New("error")
	}
	return s.tx.Transaction(ctx, updateFunc)
}

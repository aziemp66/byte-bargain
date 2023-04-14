package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB(connectionString string) *sql.DB {

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(60 * time.Minute)
	db.SetConnMaxLifetime(10 * time.Minute)

	fmt.Println("Database connected")

	return db
}

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		rollErr := tx.Rollback()
		if rollErr != nil {
			panic(rollErr)
		}
		panic(err)
	} else {
		commitErr := tx.Commit()
		if commitErr != nil {
			panic(commitErr)
		}
	}
}

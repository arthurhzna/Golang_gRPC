package database

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectDb(Ctx context.Context, connStr string) *sql.DB {

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	err = db.PingContext(Ctx)
	if err != nil {
		panic(err)
	}
	return db
}

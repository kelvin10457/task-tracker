package db

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Store struct {
	DB *sql.DB
	//sql.DB is a pool manager
	//sql.DB works with:
	//- A slice of activr connections
	//- A queue of goroutines (threads of go)
	//- Logic for destroy and create connections
}

func Open(dsn string) (*Store, error) {
	//with this you create a pool manager (a pool refers to a lot of conections)
	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(10)                  //max 10 connections at the same time
	sqlDB.SetMaxIdleConns(5)                   //max 5 connections inactive in the pool
	sqlDB.SetConnMaxLifetime(30 * time.Minute) //recicle connections every 30 min

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3) //crate a context of 3 seconds
	defer cancel()
	if err := sqlDB.PingContext(ctx); err != nil { //if in three seconds the pool don't answer you return nil and an error
		_ = sqlDB.Close()
		return nil, err
	}
	return &Store{DB: sqlDB}, nil

}

func (store *Store) Close() error {
	return store.DB.Close()
}

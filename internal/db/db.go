package db

import (
	"database/sql"
	"time"
)

type Store struct {
	DB *sql.DB
}

func Open(dsn string) (*Store, error) {
	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

}

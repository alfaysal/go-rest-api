package db

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
)

func NewMySQLStorage(cfg mysql.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	return db, nil
}

package database

import (
	"database/sql"
	"log"
	"strings"

	_ "github.com/lib/pq"
)

func InitDB(dsn string) (*sql.DB, error) {
	dsn = strings.TrimSpace(dsn)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// Test koneksi (aman untuk transaction pooler)
	// Test connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// Connection pool settings
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)

	log.Println("Database connected successfully ✅")
	return db, nil
}

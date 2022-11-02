package database

import (
	"database/sql"
	_ "github.com/lib/pq"
)

const (
	dbDriver        = "postgres"
	dbConnectionUrl = "postgresql://postgres:123123@pm-postgres:5432/pm_db?sslmode=disable"
)

type DB struct {
	driver        string
	connectionURL string
}

func New() *DB {
	return &DB{
		driver:        dbDriver,
		connectionURL: dbConnectionUrl,
	}
}

func (db *DB) Open() *sql.DB {
	dbConn, err := sql.Open(db.driver, db.connectionURL)

	if err != nil {
		panic(err)
	}

	if err := dbConn.Ping(); err != nil {
		panic(err)
	}

	return dbConn
}

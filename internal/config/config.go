package config

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/burakkarasel/IMDB-Top-250-Scrapper/internal/dsn"
)

var db *sql.DB

const maxOpenDbConn = 10
const maxIdleDbConn = 5
const maxDbLifetime = 5 * time.Minute

func Connect() error {
	d, err := sql.Open("pgx", dsn.DSN)

	if err != nil {
		log.Println("cannot open database")
		return err
	}

	err = d.Ping()

	if err != nil {
		log.Println("cannot ping database")
		return err
	}

	d.SetMaxOpenConns(maxOpenDbConn)
	d.SetConnMaxIdleTime(maxIdleDbConn)
	d.SetConnMaxLifetime(maxDbLifetime)

	db = d

	return nil
}

func GetDB() *sql.DB {
	return db
}

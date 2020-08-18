package db

import (
	"database/sql"
	"fmt"
	"log"
	"otus_final_project/src/component/config"
)

type DB struct {
	config     *config.Database
	connection *sql.DB
}

func NewDB(conf *config.Database) *DB {
	return &DB{
		config:     conf,
		connection: nil,
	}
}

func (db *DB) Connection() *sql.DB {
	if db.connection == nil {
		db.connection = db.openConnection()
	}

	return db.connection
}

func (db *DB) Close() error {
	if db.connection == nil {
		return nil
	}

	return db.connection.Close()
}

func (db *DB) openConnection() *sql.DB {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		db.config.User, db.config.Password, db.config.Host, db.config.Port, db.config.Dbname, db.config.Sslmode)

	conn, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal(err)

		return nil
	}

	return conn
}

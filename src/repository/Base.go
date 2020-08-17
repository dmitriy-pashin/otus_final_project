package repository

import (
	"database/sql"
)

type DB interface {
	Connection() *sql.DB
}

type Base struct {
	DB DB
}

func (repository Base) Connection() *sql.DB {
	return repository.DB.Connection()
}

package database

import (
	"database/sql"
)

type Database interface {
	Init() (*sql.DB, error)
}

package infrastructure

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type Database interface {
	Connect() (*sql.DB, error)
}

type PostgresConnector struct {
	Env *Env
}

func (p *PostgresConnector) Connect() (*sql.DB, error) {

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		p.Env.DBHost, p.Env.DBPort, p.Env.DBUser, p.Env.DBPass, p.Env.DBName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Successfully connected to PostgreSQL")
	return db, nil
}

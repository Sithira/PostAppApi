package infrastructure

import (
	"database/sql"
	"fmt"
)

type Application struct {
	Env *Env
}

func App() (*Application, *sql.DB) {

	app := &Application{
		Env: NewEnv(),
	}

	dbConnection, err := (&PostgresConnector{
		Env: app.Env,
	}).Connect()

	if err != nil {
		fmt.Printf("Unable to connect to database %s\n", err)
		panic("DB connection error")
	}

	return app, dbConnection
}

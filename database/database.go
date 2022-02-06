package database

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func Conectar() (*sql.DB, error) {
	conexao := "user=postgres dbname=loja_go password=password host=localhost sslmode=disable"

	db, err := sql.Open("postgres", conexao)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

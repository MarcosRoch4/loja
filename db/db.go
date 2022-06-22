package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConectDB() *sql.DB {
	conect := "user=postgres dbname=loja password=sibyl host=localhost sslmode=disable"
	db,err := sql.Open("postgres",conect) 
	
	if err != nil {
		panic(err.Error())
	}

	return db
}

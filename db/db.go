package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ConectaComBancoDeDados() *sql.DB {
	db, err := sql.Open("mysql", "victor:admin@/alura_loja")
	if err != nil {
		panic(err.Error())
	}
	return db

}

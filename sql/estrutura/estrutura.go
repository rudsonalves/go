package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	return result
}

func main() {
	db, err := sql.Open("mysql", "rudson:rafael82@/")
	if err != nil {
		panic(err)
	}
	// defer fecha o banco de dados ao final da função main
	defer db.Close()

	exec(db, "CREATE DATABASE IF NOT EXISTS cursogo")
	exec(db, "USE cursogo")
	exec(db, `CREATE TABLE IF NOT EXISTS usuarios (
		id INTEGER AUTO_INCREMENT,
		nome VARCHAR(80),
		PRIMARY KEY (id)
	)`)

}

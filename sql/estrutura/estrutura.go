package main

import (
	"database/sql"
	"fmt"

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

	fmt.Println(exec(db, "SHOW DATABASES"))
	fmt.Println(exec(db, "SELECT VERSION()"))
	exec(db, "CREATE DATABASE IF NOT EXISTS cursogo")

	// defer fecha o banco de dados ao final da função main
	defer db.Close()
}

package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "rudson:rafael82@/testdb")

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	var version string

	err = db.QueryRow("SELECT VERSION()").Scan(&version)
	if err != nil {
		panic(err)
	}

	fmt.Println(version)
	defer db.Close()
}

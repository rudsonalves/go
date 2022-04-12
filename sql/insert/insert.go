package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "rudson:rafael82@/cursogo")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmt, _ := db.Prepare("insert into usuarios(nome) values(?)")

	names := []string{"Maria", "João", "Pedro"}
	for _, name := range names {
		resp, _ := stmt.Exec(name)
		id, _ := resp.LastInsertId()
		linhas, _ := resp.RowsAffected()
		fmt.Printf("Name: %8s id: %2d linhas: %3d\n", name, id, linhas)
	}
}

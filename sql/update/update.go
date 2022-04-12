package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "rudson:rafael82@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, _ := db.Prepare("update usuarios set nome = ? where id = ?")
	stmt2, _ := db.Prepare("delete from usuarios where id = ?")

	// update
	updates := map[string]int{
		"Roberta Deangelo": 1,
		"Eduardo Peçanha":  5,
		"Flávio Souza":     8,
	}

	for nome, id := range updates {
		fmt.Println(nome, id)
		stmt.Exec(nome, id)
	}

	// delete
	stmt2.Exec(3)
}

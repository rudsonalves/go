package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "rudson:rafael82@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into usuarios(id, nome) values(?,?)")

	inserts := map[int]string{
		3:    "Bia",
		2:    "Carlos",
		4002: "José",
	}

	for id, nome := range inserts {
		_, err := stmt.Exec(id, nome)
		// este tratamento do erro é necessário para interromper a
		// execução do aplicativo
		if err != nil {
			tx.Rollback()
			log.Fatal(err)
		}
	}

	tx.Commit()
}

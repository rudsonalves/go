package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type City struct {
	Id         int
	Name       string
	Population int
}

func main() {

	db, err := sql.Open("mysql", "rudson:rafael82@/testdb")

	if err != nil {
		log.Fatal(err)
	}

	res, err := db.Query("SELECT * FROM cities")

	if err != nil {
		log.Fatal(err)
	}

	for res.Next() {

		var city City
		err := res.Scan(&city.Id, &city.Name, &city.Population)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v\n", city)
	}
	// Close all
	defer db.Close()
	defer res.Close()
}

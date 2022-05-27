package main

import (
	"fmt"
	"log"
)

type user struct {
	name string
	id   int
}

func main() {
	u := user{"alves", 1001}

	err := fmt.Errorf("usuário %s (id %d) não encontrado", u.name, u.id)

	log.Fatal(err)
}

package main

import (
	"log"
	"net/http"
)

// Executar no terminal ou será executado no diretório errado
func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

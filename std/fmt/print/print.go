package main

import "fmt"

func main() {
	nome, sobrenome := "Alberto", "Santos"
	idade := 35

	fmt.Print(nome, " ", sobrenome)
	fmt.Println(" possui", idade, "anos.")
	fmt.Println("Fim")

	fmt.Printf("%s %s possui %d anos.\nFim\n", nome, sobrenome, idade)
}

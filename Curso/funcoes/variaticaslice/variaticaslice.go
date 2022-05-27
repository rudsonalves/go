package main

import "fmt"

func imprimirAprovados(aprovados ...string) {
	fmt.Println("Lista de Aprovados:")
	for i, aprovado := range aprovados {
		fmt.Printf("aprovado[%d] %s\n", i, aprovado)
	}
}

func main() {
	aprovados := []string{"Maria", "Pedro", "Alberto", "Solange"}
	imprimirAprovados(aprovados...)
}

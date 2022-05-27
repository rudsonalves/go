package main

import "fmt"

func main() {
	nomes := []string{"Roberta", "Maria", "Alex", "Fernanda"}

	for _, nome := range nomes {
		func(n string) {
			fmt.Println("Hello " + nome)
		}(nome)
	}
}

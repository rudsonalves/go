package main

import "fmt"

/*
Forma alternatica de declarar uma função
Cria uma variável que recebe uma função anônima
*/
var soma = func(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(soma(2, 5))

	sub := func(a, b int) int {
		return a - b
	}

	fmt.Println(sub(7, 3))
}

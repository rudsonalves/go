package main

import "fmt"

func main() {
	a := []int{2, 3, 5, 7, 11, 13}

	for i, v := range a {
		fmt.Printf("Índice: %d  Valor: %d\n", i, v)
	}

	for _, v := range a {
		fmt.Printf("Valor: %d\n", v)
	}

	for i, _ := range a {
		fmt.Printf("Índice: %d\n", i)
	}
}

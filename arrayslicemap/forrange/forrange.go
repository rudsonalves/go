package main

import T "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} // copilador conta os elementos da array

	T.Println(numeros)

	for i, numero := range numeros {
		T.Printf("numeros[%d] %d\n", i, numero)
	}

	// pegando apenas o conteúdo da array
	for _, numero := range numeros {
		T.Printf("numeros[x] %d\n", numero)
	}

	/*
		Cuidado!!!
		for i, numero := range numeros {
			T.Printf("numeros[x] %d\n", numero)
		}

		Este código dará erro porque a variável i, declarada, não é usada
	*/
}

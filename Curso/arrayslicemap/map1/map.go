package main

import T "fmt"

func main() {
	// var aprovados map][int]string
	// maps devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[123456788] = "Maria"
	aprovados[434525653] = "Pedro"
	aprovados[634345675] = "Ana"

	for cpf, nome := range aprovados {
		T.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	T.Printf("\nRemovido CPF 434525653:\n")
	delete(aprovados, 434525653)

	for cpf, nome := range aprovados {
		T.Printf("%s (CPF: %d)\n", nome, cpf)
	}
}

package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva":   12323.34,
			"Gilberto Fonseca": 34234.34,
		},
		"P": {
			"Paulo Alberto":   234435.23,
			"Patrick Almeida": 654464.23,
		},
		"J": {
			"José Maria Feurosa": 334553.34,
		},
	}

	// delete(funcsPorLetra, "P")
	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}

	for _, funcs := range funcsPorLetra {
		for nome, salario := range funcs {
			fmt.Printf("Nome: %s  Salario: %.2f\n", nome, salario)
		}
	}
	fmt.Println("...")
	delete(funcsPorLetra, "P")
	for _, funcs := range funcsPorLetra {
		for nome, salario := range funcs {
			fmt.Printf("Nome: %s  Salario: %.2f\n", nome, salario)
		}
	}
}

package main

import "fmt"

func main() {
	funcESalarios := map[string]float64{
		"José João":         12123.34,
		"Gabriela Silva":    323423.34,
		"Fernanda Nogueira": 357633.24,
	}

	for nome, salario := range funcESalarios {
		fmt.Printf("%s: R$%.2f\n", nome, salario)
	}

	delete(funcESalarios, "Alberto Nunes")
	// remover um elemento inixistente não gera erro
}

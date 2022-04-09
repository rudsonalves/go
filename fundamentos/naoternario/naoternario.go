package main

import "fmt"

// Não tem operador ternário
func obterResultado(nota float64) string {
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"
	/* Poderia ser veito assim:
	if nota >= 6 {
		return "Aprovado"
	} else {
		return "Reprovado"
	}
	*/
}

func main() {
	fmt.Println("5.6: ", obterResultado(5.6))
	fmt.Println("6.2: ", obterResultado(6.2))
}

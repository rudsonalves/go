package main

import T "fmt"

func imprimirResultado(nota float64) {
	if nota >= 7 {
		T.Printf("Aprovado com nota %.1f\n", nota)
	} else {
		T.Printf("Reprovado com nota %.1f\n", nota)
	}
}

func main() {
	imprimirResultado(7.3)
	imprimirResultado(5.234)
	imprimirResultado(6.995)
}

package main

import "fmt"

func notaParaConceito(n float64) string {
	if n >= 9 && n <= 10 {
		return "A"
	} else if n >= 8 && n < 9 {
		return "B"
	} else if n >= 5 && n < 8 {
		return "C"
	} else if n >= 3 && n < 5 {
		return "D"
	} else {
		return "E"
	}
}

func imprimiNota(n float64) {
	fmt.Printf("Nota %.1f: conceito %s\n", n, notaParaConceito(n))
}

func main() {
	imprimiNota(9.8)
	imprimiNota(8.2)
	imprimiNota(6.9)
	imprimiNota(4.5)
	imprimiNota(1.3)
}

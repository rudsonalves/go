package main

import "fmt"

func notaParaConceito(n float64) string {
	switch {
	case n >= 9 && n <= 10:
		return "A"
	case n >= 8 && n < 9:
		return "B"
	case n >= 5 && n < 8:
		return "C"
	case n >= 3 && n < 5:
		return "D"
	default:
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

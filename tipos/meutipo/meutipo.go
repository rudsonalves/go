package main

import "fmt"

type nota float64

func (n nota) entre(inicio, fim float64) bool {
	return float64(n) >= inicio && float64(n) <= fim
	// o float64() aqui é necessário para converter o nota em float64
}

func notaParaConceito(n nota) string {
	if n.entre(9., 10.) {
		return "A"
	} else if n.entre(7.0, 8.99) {
		return "B"
	} else if n.entre(5.0, 6.99) {
		return "C"
	} else if n.entre(3.0, 4.99) {
		return "D"
	}
	return "E"
}

func main() {
	fmt.Println(notaParaConceito(7.6))
	fmt.Println(notaParaConceito(9.4))
	fmt.Println(notaParaConceito(5.1))
	fmt.Println(notaParaConceito(4.2))
	fmt.Println(notaParaConceito(1.3))
}

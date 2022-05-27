package main

import "fmt"

func fatorial(n int) (int, error) {
	switch {
	case n < 0:
		return -1, fmt.Errorf("Número inválido: %d", n)
	case n == 0:
		return 1, nil
	default:
		fatorialAnterior, _ := fatorial(n - 1)
		return n * fatorialAnterior, nil
	}
}

func fatorial2(n uint) uint {
	if n == 0 {
		return 1
	}

	fatorialAnterior := fatorial2(n - 1)
	return n * fatorialAnterior
}

func main() {
	for i := uint(0); i < 11; i++ {
		resultado := fatorial2(i)
		fmt.Printf("Fatorial de %2d é: %10d\n", i, resultado)
	}
}

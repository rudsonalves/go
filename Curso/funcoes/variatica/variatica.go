package main

import "fmt"

func media(numeros ...float64) float64 {
	soma := 0.0
	for _, num := range numeros {
		soma += num
	}
	return soma / float64(len(numeros))
}

func main() {
	fmt.Printf("Média: %.2f", media(2.3, 4.8, 5.2, 8.9, 2.5, 9.5))
}

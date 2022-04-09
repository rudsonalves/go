package main

import T "fmt"

func main() {
	var notas [3]float64
	T.Println(notas)

	notas[0], notas[1], notas[2] = 7.8, 8.2, 5.7

	T.Println(notas)

	total := 0.0
	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}

	media := total / float64(len(notas))

	T.Printf("Média: %.2f\n", media)
}

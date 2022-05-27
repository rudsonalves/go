package main

import (
	"fmt"
	"strconv"
)

func main() {
	const (
		msg1 = "\nO número %d é divisível por %d e por isto não é primo\n\n"
		msg2 = "\nParabéns \"%d\" é primo.\n\n"
	)

	var (
		nstr          string = ""
		primos               = []int{2, 3, 5, 7, 11}
		numero, index int    = 0, 0
	)

	for numero == 0 {
		fmt.Print("\nEntre com um primo menor que 100: ")
		fmt.Scanln(&nstr)
		n, err := strconv.Atoi(nstr)
		if err != nil {
			fmt.Printf("\n%q não é um numero\n", nstr)
			continue
		} else if (n > 0) && (n <= 100) {
			numero = n
			continue
		}
		fmt.Printf("\n%q é maior que 100!\n", nstr)
	}

	for index < len(primos) {
		if numero%primos[index] == 0 {
			fmt.Printf(msg1, numero, primos[index])
			index = -1
			break
		}
		index++
	}

	if index >= 0 {
		fmt.Printf(msg2, numero)
	}
}

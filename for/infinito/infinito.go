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
		nstr            string = ""
		primos                 = []int{2, 3, 5, 7, 11}
		numero, e_primo        = 0, true
	)

	for {
		fmt.Print("\nEntre com um primo menor que 100: ")
		fmt.Scanln(&nstr)
		n, err := strconv.Atoi(nstr)
		if err != nil {
			fmt.Printf("\n%q não é um numero\n", nstr)
			continue
		} else if (n > 0) && (n <= 100) {
			numero = n
			break
		}
		fmt.Printf("\n%q é maior que 100!\n", nstr)
	}

	for _, p := range primos {
		if numero%p == 0 {
			fmt.Printf(msg1, numero, p)
			e_primo = false
			break
		}
	}

	if e_primo {
		fmt.Printf(msg2, numero)
	}
}

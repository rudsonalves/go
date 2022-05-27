package main

import "fmt"

func main() {
	var a, b, c, m int

	for {
		fmt.Println("\nEnter com os valores de a b c<ENTER>")
		n, err := fmt.Scanln(&a, &b, &c)
		if err == nil {
			m = n
			break
		}
		fmt.Printf("Entrou %d e esperado 3 entradas. Erro:%v\n", n, err)
	}

	fmt.Printf("Entrou com %d números, de valores: a = %d, b = %d e c = %d\n", m, a, b, c)
}

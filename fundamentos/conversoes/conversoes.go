package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4 // float
	y := 2   // inteiro

	fmt.Println(x / float64(y))
	/* x/y gera erro pela minstura de tipos
	(mismatched types float64 and int) */

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	// cuidado na conversão para string ...
	fmt.Println("Teste " + string(97))
	// converte para o caractere "a" e não a stroing 97
	fmt.Println("Novamento com o valor inteiro " + strconv.Itoa(97))

	// mais sobre o strconv: ele retorna dois valores
	n1, e1 := strconv.Atoi("123")

	fmt.Println(n1, e1)

	// ignorando o erro
	n2, _ := strconv.Atoi("123")
	fmt.Println(n2)

	// booleano
	b, _ := strconv.ParseBool("true")
	if b {
		fmt.Println("Verdadeiro")
	}
}

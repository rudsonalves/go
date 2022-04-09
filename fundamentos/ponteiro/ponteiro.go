package main

import "fmt"

func main() {
	i := 1

	// Go não tem aritimética de ponteiros mas pode referenciar
	var p *int = nil

	p = &i // pegando o endereço de i

	*p++
	i++

	// p++ seria uma aritimética de ponteiro e não pode ser feito

	fmt.Println(p, *p, i, &i)
}

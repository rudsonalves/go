package main

import (
	"fmt"
)

type binario byte

const (
	_ = 1 << (10 * iota)
	KiB
	MiB
	GiB
	TiB
	PiB
	EiB
)

const (
	domingo = iota + 1
	segunda
	terca
	quarta
	quinta
	sexta
	sabado
)

const (
	micro = 3*iota - 6
	mili
	_
	quilo
	mega
	giga
)

func main() {
	fmt.Println("Potências Binárias:")
	fmt.Println("KiB: ", KiB)
	fmt.Println("MiB: ", MiB)
	fmt.Println("GiB: ", GiB)
	fmt.Println("TiB: ", TiB)
	fmt.Println("PiB: ", PiB)
	fmt.Println("EiB: ", EiB)

	fmt.Println("\nDias da semana:")
	fmt.Println("D: ", domingo)
	fmt.Println("S: ", segunda)
	fmt.Println("S: ", terca)
	fmt.Println("S: ", quarta)
	fmt.Println("S: ", quinta)
	fmt.Println("S: ", sexta)
	fmt.Println("S: ", sabado)

	fmt.Println("\nPrefixos:")
	fmt.Println("u: 10^", micro)
	fmt.Println("m: 10^", mili)
	fmt.Println("K: 10^", quilo)
	fmt.Println("M: 10^", mega)
	fmt.Println("G: 10^", giga)
}

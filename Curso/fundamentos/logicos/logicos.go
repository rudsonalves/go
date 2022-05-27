package main

import "fmt"

func compras(trab1, trab2 bool) (bool, bool, bool) {
	comprarTV50 := trab1 && trab2
	compratTV32 := trab1 != trab2
	comprarSorvete := trab1 || trab2

	return comprarTV50, compratTV32, comprarSorvete
}

func main() {
	tv50, tv32, sorvete := compras(true, true)
	fmt.Printf("Tv50: %t, Tv32: %t, Sorvete: %t, Saudável: %t", tv50, tv32, sorvete, !sorvete)

	tv50, tv32, sorvete = compras(true, false)
	fmt.Printf("\nTv50: %t, Tv32: %t, Sorvete: %t, Saudável: %t", tv50, tv32, sorvete, !sorvete)

	tv50, tv32, sorvete = compras(false, false)
	fmt.Printf("\nTv50: %t, Tv32: %t, Sorvete: %t, Saudável: %t", tv50, tv32, sorvete, !sorvete)

}

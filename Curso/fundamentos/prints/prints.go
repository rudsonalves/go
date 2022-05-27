package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Println(" Nova")
	fmt.Println("Linha")

	x := 3.141516

	/* fmt.Println("O valor de x é" + x)
	concatenar conteúdos */
	xs := fmt.Sprint(x)

	fmt.Println("O valor de x é " + xs)

	fmt.Printf("O valor de x é %.2f.", x)

	a, b, c, d := 1, 1.9999, false, "opa"

	fmt.Printf("\n%d %f %.2f %t %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %v %v", a, b, c, d)
}

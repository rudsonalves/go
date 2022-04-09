package main

import "fmt"

func main() {
	var b byte = 3
	fmt.Println(b)

	i := 3 // inferencia de tipo
	i += 3
	fmt.Println(i)
	i -= 3
	fmt.Println(i)
	i /= 3
	fmt.Println(i)
	i *= 5
	fmt.Println(i)
	i += 5
	fmt.Println(i)
	i %= 3
	fmt.Println(i)

	x, y := 1, 2 // cria as variáveis e inicia
	fmt.Printf("x = %d, y = %d", x, y)
	// Permuta os valores das variáveis
	x, y = y, x // usar apenas o '=' pois já foram criadas anteriormente
	fmt.Printf("\nx = %d, y = %d", x, y)
}

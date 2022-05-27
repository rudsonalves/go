package main

import "fmt"

// Valores zeros atribuídos por padrão
func main() {
	var a int
	var b float64
	var c bool
	var d string
	var e *int // ponteiro para um inteiro

	fmt.Printf("%v %v %v %q %v", a, b, c, d, e)
}

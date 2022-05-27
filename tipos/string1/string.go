package main

import (
	"fmt"
)

func main() {
	frase := `"Duas coisas são infinitas: o universo e a estupidez 
humana. Mas, em relação ao universo, ainda não tenho certeza 
absoluta."

Albert Einstein

`

	letra := 'a'
	escape := "O importante é não parar de 'questionar'."

	fmt.Println(frase)
	fmt.Printf("%v %T\n", letra, letra)
	fmt.Println(escape)
}

package main

import (
	"fmt"
)

func main() {
	populacao := map[string]int{
		"Brasil":         213317639,
		"China":          1411780000,
		"Índia":          1380004385,
		"Estados Unidos": 331449281,
		"Indonésia":      273523615,
		"Paquistão":      220892340,
	}

	pop, ok := populacao["Brasil"]
	fmt.Println(pop, ok)

	pop, ok = populacao["Espanha"]
	fmt.Println(pop, ok)

	pop = populacao["China"]
	fmt.Println(pop)
}

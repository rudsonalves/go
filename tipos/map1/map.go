package main

import "fmt"

func main() {
	populacao := map[string]int{}
	// var populacao map[string]int64

	populacao["Brasil"] = 213317639
	populacao["China"] = 1411780000
	populacao["Índia"] = 1380004385
	populacao["Estados Unidos"] = 331449281
	populacao["Indonésia"] = 273523615
	populacao["Paquistão"] = 220892340

	// for pais, pop := range populacao {
	// 	fmt.Printf("Pais: %15s  %10d\n", pais, pop)
	// }

	fmt.Println(populacao)
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	nome := "Alberto Santos Dumont"

	split := strings.Fields(nome)

	fmt.Println(split)

	join := strings.Join(split, "_")

	fmt.Println(join)
}

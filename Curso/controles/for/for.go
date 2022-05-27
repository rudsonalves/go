package main

import (
	"fmt"
	"time"
)

func main() {
	// for como while
	fmt.Println("For como while:")
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	fmt.Println("\nFor convencional:")
	for i := 0; i <= 20; i += 2 {
		fmt.Println(i)
	}

	fmt.Println("\nMisturando:")
	for i := 1; i < 10; i++ {
		if i%2 == 0 {
			fmt.Printf("%d é par\n", i)
		} else {
			fmt.Printf("%d é impar\n", i)
		}
	}

	fmt.Println("\nLaço infinito (parar com CTRL+ALT+M):")
	for {
		fmt.Print(".")
		time.Sleep(time.Second * 2)
	}
}

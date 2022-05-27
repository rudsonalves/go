package main

import "fmt"

func main() {
	fmt.Println()

	var hello = func() {
		fmt.Print("Hello World!!!")
	}

	hello()
}

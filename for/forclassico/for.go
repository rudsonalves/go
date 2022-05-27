package main

import "fmt"

func main() {
	fmt.Println(" D   S   T   Q   Q   S   S")
	for i := 0; i < 31; i += 7 {
		for j := 0; j < 7; j++ {
			dia := i + j + 1
			fmt.Printf("%2d  ", dia)
		}
		fmt.Println()
	}
}

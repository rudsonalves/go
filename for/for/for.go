package main

import "fmt"

func main() {
	fmt.Println(" D   S   T   Q   Q   S   S")
	for dia := 1; dia < 31; dia += 7 {
		for j := 0; j < 7; j++ {
			fmt.Printf("%2d  ", dia+j)
		}
		fmt.Println()
	}
}

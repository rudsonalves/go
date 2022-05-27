package main

import (
	"fmt"
)

var start = 0

func init() {
	fmt.Println("This is init()...")
	start = 10
}

func main() {
	fmt.Println("This is main()...")
	fmt.Println("Start:", start)
}

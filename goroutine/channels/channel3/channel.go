package main

import (
	"fmt"
)

func print(id int, c <-chan int) {
	fmt.Printf("(%d) channel: %v\n", id, <-c)
}

func main() {
	ch := make(chan int)

	go print(1, ch)
	go print(2, ch)
	go print(3, ch)

	ch <- 0
	ch <- 1
	ch <- 2
}

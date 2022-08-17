package main

import "fmt"

func multipliByTwo(in chan int, out chan string, id int) {
	fmt.Printf("Inicializing go routine (%d)...\n", id)
	for {
		num := <-in
		result := num * 2
		out <- fmt.Sprintf("by (%d): %d", id, result)
	}
}

func main() {
	out := make(chan string)
	in := make(chan int, 4)

	go multipliByTwo(in, out, 1)
	go multipliByTwo(in, out, 2)
	go multipliByTwo(in, out, 3)

	//go func() {
	in <- 1
	in <- 2
	in <- 3
	in <- 4
	//}()

	fmt.Println(<-out)
	fmt.Println(<-out)
	fmt.Println(<-out)
	fmt.Println(<-out)
}

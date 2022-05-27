package main

import (
	"fmt"
)

func isPrimo(num int) bool {
	for i := 2; i < num/2+1; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func primos(n int, c chan int) {
	inicio := 2
	for i := 0; i < n; i++ {
		for primo := inicio; ; primo++ {
			if isPrimo(primo) {
				c <- primo
				inicio = primo + 1
				// time.Sleep(time.Microsecond * 100)
				break
			}
		}
	}
	close(c)
}

func main() {
	c := make(chan int, 30)
	go primos(100, c)
	for primo := range c {
		fmt.Printf("%d ", primo)
	}
	fmt.Println("\n Fim!")
}

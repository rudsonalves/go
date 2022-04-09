package main

import "fmt"

func eprimo(n int) bool {
	if n <= 3 {
		return true
	}

	for i := 2; i < n/2+1; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	for i := 1; i < 100; i++ {
		if eprimo(i) {
			fmt.Println(i, "é primo")
			// } else {
			// 	fmt.Println(i, "não é primo")
		}
	}
}

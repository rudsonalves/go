package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if (i%3 == 0) && (i%5 == 0) {
			fmt.Println("5 and 3:", i)
		} else if i%3 == 0 {
			fmt.Println("3:", i)
		} else if i%5 == 0 {
			fmt.Println("5:", i)
			continue
		} else {
			fmt.Println(i)
		}
	}
}

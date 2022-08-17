package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Use primes <number>")
		os.Exit(0)
	}
	n, ok := strconv.Atoi(os.Args[1])
	if ok != nil {
		n = 5
	}

	var primes = []int{2, 3}
	n--
	if n > 3 {
		i := 3
		for len(primes) <= n {
			i += 2
			j := int(math.Sqrt(float64(i)))
			for _, p := range primes {
				if p > j {
					// fmt.Printf("%d é primo.\n", i)
					primes = append(primes, i)
					break
				}
				if i%p == 0 {
					// fmt.Printf("%d não é primo. Divisível por %d.\n", i, p)
					break
				}
			}
		}
	}

	fmt.Println(primes)
}

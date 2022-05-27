package main

import "fmt"

func main() {

	var primos = map[int]bool{}
	max_index := 200

	for i := 2; i < max_index; i++ {
		primos[i] = true
	}

	mult_2, mult_3, mult_5, mult_7, mult_11, mult_13, mult_17 := 0, 0, 0, 0, 0, 0, 0
	for i := 1; i < max_index; i++ {
		index := i * 2
		if index < max_index {
			if primos[index] && index != 2 {
				delete(primos, index)
				mult_2++
			}
		}

		index = i * 3
		if index < max_index {
			if primos[index] && index != 3 {
				delete(primos, index)
				mult_3++
			}
		}

		index = i * 5
		if index < max_index {
			if primos[index] && index != 5 {
				delete(primos, index)
				mult_5++
			}
		}

		index = i * 7
		if index < max_index {
			if primos[index] && index != 7 {
				delete(primos, index)
				mult_7++
			}
		}

		index = i * 11
		if index < max_index {
			if primos[index] && index != 11 {
				delete(primos, index)
				mult_11++
			}
		}

		index = i * 13
		if index < max_index {
			if primos[index] && index != 13 {
				delete(primos, index)
				mult_13++
			}
		}

		index = i * 17
		if index < max_index {
			if primos[index] && index != 17 {
				delete(primos, index)
				mult_17++
			}
		}
	}

	sortPrimos := []int{}
	for i := 1; i < max_index; i++ {
		if primos[i] {
			sortPrimos = append(sortPrimos, i)
		}
	}
	fmt.Println(sortPrimos, "\nTotal de primos:", len(sortPrimos))
	fmt.Println("\nMultiplos de  2:", mult_2)
	fmt.Println("Multiplos de  3:", mult_3)
	fmt.Println("Multiplos de  5:", mult_5)
	fmt.Println("Multiplos de  7:", mult_7)
	fmt.Println("Multiplos de 11:", mult_11)
	fmt.Println("Multiplos de 13:", mult_13)
	fmt.Println("Multiplos de 17:", mult_17)
}

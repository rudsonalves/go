package main

import "fmt"

func haveNumber(numbers []int, num int) (bool, int) {
	for index, n := range numbers {
		if n == num {
			return true, index
		}
	}
	return false, 0
}

func removeIndex(n []int, index int) []int {
	newNumbers := make([]int, len(n)-1)
	if index == 0 {
		copy(newNumbers, n[1:])
	} else if index == len(n)-1 {
		copy(newNumbers, n[:len(n)-1])
	} else {
		newNumbers = make([]int, index)
		copy(newNumbers, n[:index])
		newNumbers = append(newNumbers, n[index+1:]...)
	}
	return newNumbers
}

func getMaxPrimo(primos []int, maxNum int) int {
	for id, v := range primos {
		if v*v > maxNum {
			return primos[id-1]
		}
	}
	return 99
}

func main() {

	var numbers = []int{}
	max_index := 100

	max_index++
	for i := 2; i < max_index; i++ {
		numbers = append(numbers, i)
	}

	// fmt.Println(numbers)
	var primos = []int{2}

	pIndex := 0
	for {
		p := primos[pIndex]
		p2 := p * p
		if p2 > max_index {
			break
		}

		// fmt.Println("Remove multiples of", p)
		n := 0
		for i := p; ; i++ {
			remove := i * p
			if remove > max_index {
				break
			}
			ok, index := haveNumber(numbers, remove)
			if remove < max_index && ok {
				numbers = removeIndex(numbers, index)
				n++
				// fmt.Println(numbers)
			}
		}
		// fmt.Printf("Removed %4d multiples of %d\n", n, p)

		pIndex++
		if pIndex >= len(primos) {
			for _, v := range numbers[pIndex:] {
				if v > p2 {
					break
				}
				primos = append(primos, v)
			}
			// fmt.Println(primos)
		}
	}
	fmt.Println(numbers)
}

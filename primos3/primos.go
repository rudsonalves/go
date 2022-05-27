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

func main() {

	var prime = []int{2}
	max_index := 1000000

	max_index++
	for i := 3; i < max_index; i += 2 {
		prime = append(prime, i)
	}

	pIndex := 1
	for {
		p := prime[pIndex]
		p2 := p * p
		if p2 > max_index {
			break
		}

		for i := p; ; i += 2 {
			remove := i * p
			if remove > max_index {
				break
			}
			ok, index := haveNumber(prime, remove)
			if remove < max_index && ok {
				prime = removeIndex(prime, index)
			}
		}
		pIndex++
	}
	fmt.Println(prime)
}

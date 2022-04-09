package main

import (
	"errors"
	"fmt"
)

// func div(dividendo, divisor int) (divisao, resto int, err error) {
// 	if divisor != 0 {
// 		divisao = dividendo / divisor
// 		resto = dividendo % divisor
// 		err = nil
// 		return
// 	}
// 	divisao = 0
// 	resto = 0
// 	err = errors.New("Can't divide by zero")
// 	return
// }

func div(dividendo, divisor int) (int, int, error) {
	if divisor != 0 {
		return dividendo / divisor, dividendo % divisor, nil
	}
	return 0, 0, errors.New("Can't divide by zero")
}

func main() {
	d, r, _ := div(13, 4)
	fmt.Printf("Divisão 13/4: %d; resto: %d", d, r)
}

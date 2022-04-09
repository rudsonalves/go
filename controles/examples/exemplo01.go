package main

import "fmt"

func main() {
	evenVals := []int{2, 4, 6, 8, 10}
	for i, v := range evenVals {
		if i == 0 {
			continue
		}
		if i == len(evenVals)-2 {
			break
		}
		fmt.Println(i, v)
	}

	fmt.Println(">>")
	for i := 1; i < len(evenVals)-1; i++ {
		fmt.Println(i, evenVals[i])
	}

}

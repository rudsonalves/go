package main

import "fmt"

func main() {
	var x [3]int
	var y = [3]int{1, 2, 3}
	var z = [12]int{1, 5: 4, 6, 10: 100, 15}
	var x1 = [3]int{1, 2, 3}
	var m [2][3]int

	fmt.Println(x, y, z)
	fmt.Println(y == x1)
	fmt.Println(m)
	fmt.Println(len(m[1]))
}

package main

import "fmt"

func main() {
	v := [4]int{}
	z := [3]int{2, 4, 6}
	w := [10]int{1, 1, 5: 4, 8: 7}
	m := [3][2]int{{1, 2}, {3, 4}, {5, 6}}

	// fmt.Println("v: ", v)
	// fmt.Println("z: ", z)
	// fmt.Println("w: ", w)
	// fmt.Println("m: ", m)

	v[0] = z[0] + w[8]    // 2 + 7 = 9
	v[1] = w[1] + m[2][1] // 1 + 6 = 7
	v[2] = m[1][1]        // 4
	v[3] = 12

	fmt.Println("v:", v, "Comprimento:", len(v))
}

package main

import "fmt"

func main() {
	var s string = "Hello there"
	var b byte = s[6]
	s0 := s[2:4]
	fmt.Printf("%q %d %q len(s):%d\n", s, b, s0, len(s))

	var a rune = 'x'
	var s1 string = string(a)
	var b1 byte = 'y'
	var s2 string = string(b1)
	fmt.Println(" a:", a, "s1:", s1)
	fmt.Println("b1:", b1, "s2:", s2)
}

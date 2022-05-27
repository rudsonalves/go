package main

import "fmt"

type Adder struct {
	start int
}

func (a Adder) AddTo(val int) int {
	return a.start + val
}

func main() {
	myAdder := Adder{10}
	fmt.Println(myAdder.AddTo(5))

	f1 := myAdder.AddTo
	fmt.Println(f1(10))

	f2 := myAdder.AddTo
	fmt.Println(f2(15))
}

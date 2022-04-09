package main

import "fmt"

type person struct {
	name string
	age  int
	pet  string
}

func main() {
	var fred person

	fred.name = "Fred Alborg"
	fred.age = 24
	fred.pet = "Doug"

	julia := person{
		"Julia Robson",
		40,
		"cat",
	}

	fmt.Println("Name:", fred.name, " Age:", fred.age, " Pet:", fred.pet)
	fmt.Println(julia)
}

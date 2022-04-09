package main

import (
	T "fmt"
	"time"
)

func main() {
	T.Println("Strings:", "Banana" == "Banana")
	T.Println("!=", 3 != 2)
	T.Println("Outros operadores <, >, <=, >=")

	// Datas
	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)
	T.Println("Datas:", d1 == d2)
	T.Println("Datas:", d1.Equal(d2))

	// Struct
	type Pessoa struct {
		Nome string
	}

	p1 := Pessoa{"João"}
	p2 := Pessoa{"João"}
	p3 := Pessoa{"Maria"}
	T.Println("Pessoa p1 == p2:", p1 == p2)
	T.Println("Pessoa p1 == p3:", p1 == p3)

}

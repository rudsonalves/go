package main

import "fmt"

func f1() {
	fmt.Println("F1")
}

func f2(p1 string, p2 string) {
	fmt.Printf("F2: %s %s\n", p1, p2)
}

func f3() string {
	return "F3"
}

func f4(p1, p2 string) string {
	return fmt.Sprintf("F4: %s %s\n", p1, p2)
}

func f5() (string, string) {
	return "Parametro 1", "Parametro 2"
}

func main() {
	f1()
	f2("Rudson", "Alves")

	r3, r4 := f3(), f4("Juliana", "Boasq")

	fmt.Println(r3)
	fmt.Println(r4)

	r5a, r5b := f5()
	fmt.Printf("F5: %s, %s\n", r5a, r5b)
	_, r5 := f5()
	fmt.Printf("F5: %s\n", r5)
}

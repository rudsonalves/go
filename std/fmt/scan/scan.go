package main

import "fmt"

func main() {
	var a, b, c, d int64

	fmt.Println("Entre com quatro valores: ")
	fmt.Scan(&a, &b, &c, &d)
	fmt.Printf("a:%d  b:%d  c:%d  d:%d\n", a, b, c, d)

	fmt.Println("\nEntre com três valores na forma %%d %%d<ENTER> %%d %%d<ENTER>: ")
	fmt.Scanf("%d %d\n %d %d\n", &a, &b, &c, &d)
	fmt.Printf("a:%d  b:%d  c:%d  d:%d\n", a, b, c, d)

	fmt.Println("\nEntre com quatro valores na forma com um \\n: ")
	fmt.Scanln(&a, &b, &c, &d)
	fmt.Printf("a:%d  b:%d  c:%d  d:%d\n", a, b, c, d)
}

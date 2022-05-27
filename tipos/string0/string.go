package main

import (
	"fmt"
)

func main() {
	for i := 127744; i < 127891; i++ {
		fmt.Printf("%d %3s   ", i, string(rune(i)))
		if (i+2)%5 == 0 {
			fmt.Println("")
		}
	}

}

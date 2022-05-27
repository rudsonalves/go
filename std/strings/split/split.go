package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "a,b,c,d"

	fmt.Printf("      Split: %q\n", strings.Split(str, ","))
	fmt.Printf("     SplitN: %q\n", strings.SplitN(str, ",", 2))
	fmt.Printf(" SplitAfter: %q\n", strings.SplitAfter(str, ","))
	fmt.Printf("SplitAfterN: %q\n", strings.SplitAfterN(str, ",", 2))
}

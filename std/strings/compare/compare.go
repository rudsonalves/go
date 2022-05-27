package main

import (
	"fmt"
	"strings"
)

func main() {
	b, a, f := strings.Cut("Albert Eintein", "ert")
	fmt.Printf("before: %q\n after: %q\n found: %t \n\n", b, a, f)

	b, a, f = strings.Cut("Albert Eintein", "e")
	fmt.Printf("before: %q\n after: %q\n found: %t \n\n", b, a, f)

	b, a, f = strings.Cut("Albert Eintein", "see")
	fmt.Printf("before: %q\n after: %q\n found: %t \n\n", b, a, f)
}

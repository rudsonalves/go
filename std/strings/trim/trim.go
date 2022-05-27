package main

import (
	"fmt"
	"strings"
	"unicode"
)

func removeSelection(r rune) bool {
	if unicode.IsNumber(r) || r == ' ' || r == 'o' {
		return true
	}
	return false
}

func main() {
	var (
		str1 = "---Mais sem graça que a top-model ..."
		str2 = "  Mas ontem eu recebi um telegrama   "
		str3 = "o Prêmio Nobel de Física de 1921"
	)

	fmt.Printf("%q\n", strings.Trim(str1, "- ."))
	fmt.Printf("%q\n", strings.TrimLeft(str1, "- ."))
	fmt.Printf("%q\n", strings.TrimRight(str1, "- ."))
	fmt.Printf("%q\n", strings.TrimPrefix(str1, "-"))
	fmt.Printf("%q\n", strings.TrimSuffix(str1, "."))
	fmt.Printf("%q\n", strings.TrimSpace(str2))
	fmt.Printf("%q\n", strings.TrimRightFunc(str3, unicode.IsNumber))
	fmt.Printf("%q\n", strings.TrimFunc(str3, removeSelection))
}

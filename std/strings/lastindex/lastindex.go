package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	str := "A Revolução Americana, ou Revolução Americana de 1776, teve suas raízes"

	fmt.Println(strings.LastIndexFunc(str, unicode.IsNumber))
	fmt.Println(strings.LastIndexFunc(str, unicode.IsPunct))
	fmt.Println(strings.LastIndexFunc(str, unicode.IsLetter))
	fmt.Println(strings.LastIndexFunc(str, unicode.IsSpace))
	fmt.Println(strings.LastIndexFunc(str, unicode.IsUpper))
}

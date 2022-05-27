package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "A Revolução Americana, ou Revolução Americana de 1776, teve suas raízes"

	fmt.Println(strings.Replace(str, "Americana", "___________", 2))
	fmt.Println(strings.Replace(str, "Americana", "___________", -1))
	fmt.Println(strings.ReplaceAll(str, "Americana", "___________"))
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "prIncíPios básICos do paSTafArismo"

	fmt.Println(strings.Title(str))
	fmt.Println(strings.ToTitle(str))
	fmt.Println(strings.ToUpper(str))
	fmt.Println(strings.ToLower(str))
	fmt.Println(strings.Title(strings.ToLower(str)))
}

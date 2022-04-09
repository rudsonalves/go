package main

import "fmt"

func main() {
	words := []string{"a", "cow", "smile", "gopher", "octopus", "anthropologist"}

	for _, word := range words {
		// fmt.Println(word, len(word))

		// switch wordLen := len(word); wordLen {
		// case 1, 2, 3, 4:
		// 	fmt.Printf("%q is a short word\n", word)
		// case 5:
		// 	fmt.Printf("%q is exactly the right legth: %d\n", word, len(word))
		// case 6, 7, 8, 9:
		// default:
		// 	fmt.Printf("%q is a long word\n", word)
		// }

		// In this code the sentence "word Len := len(world);" is to set 
		// the wordLen variable. Since wordLen is not declared as a key in 
		// the switch statement, it is necessary to invoke wordLen in case 
		// sentences
		switch wordLen := len(word); {
		case wordLen < 5:
			fmt.Printf("%q is a short word\n", word)
		case wordLen == 5:
			fmt.Printf("%q is exactly the right legth: %d\n", word, wordLen)
		case wordLen < 10:
		default:
			fmt.Printf("%q is a long word\n", word)
		}
	}
}

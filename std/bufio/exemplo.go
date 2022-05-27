package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Seu nome completo:")

	scanner := bufio.NewReader(os.Stdin)
	line, _ := scanner.ReadString('\n')

	fmt.Printf("%q %T\n", line, line)
}

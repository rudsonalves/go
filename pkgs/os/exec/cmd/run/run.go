package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmd1 := exec.Command("ls", "-lah", "/usr")
	cmd2 := exec.Command("grep", "lib")

	b1 := bytes.Buffer{}
	b2 := bytes.Buffer{}

	cmd1.Stdout = &b1
	cmd2.Stdin = &b1
	cmd2.Stdout = &b2

	if err := cmd1.Run(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s:\n%s\n\n", cmd1.String(), b1.String())

	if err := cmd2.Run(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf(" ...| %s:\n%s\n", cmd2.String(), b2.String())
}

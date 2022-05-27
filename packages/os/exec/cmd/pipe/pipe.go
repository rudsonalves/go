package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmd1 := exec.Command("lsmod")
	stdout1 := bytes.Buffer{}
	cmd1.Stdout = &stdout1

	cmd2 := exec.Command("grep", "^sound")
	stdout2 := bytes.Buffer{}
	cmd2.Stdin = &stdout1
	cmd2.Stdout = &stdout2

	if err := cmd1.Run(); err != nil {
		log.Fatal(err)
	}

	if err := cmd2.Run(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(stdout2.String())
}

package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	home := os.Getenv("HOME")
	cmd := exec.Command("ls", "-laRh", home)

	buffer := bytes.Buffer{}
	cmd.Stdout = &buffer

	if err := cmd.Start(); err != nil {
		log.Fatal()
	}

	fmt.Printf("Executando um %s...\n", cmd.String())
	cmd.Wait()
	fmt.Println("Terminou...")
}

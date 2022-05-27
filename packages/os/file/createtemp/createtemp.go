package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.CreateTemp("", "exemple.*.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer os.Remove(file.Name())

	if _, err := file.Write([]byte("content")); err != nil {
		file.Close()
		log.Fatal(err)
	}

	fmt.Println("Name: ", file.Name())
	stat, _ := file.Stat()
	fmt.Println("Size: ", stat.Size())
	fmt.Println("Time: ", stat.ModTime())

	if err := file.Close(); err != nil {
		log.Fatal(err)
	}
}

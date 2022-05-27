package main

import (
	"fmt"
	"log"
	"os"
)

func IsFile(d os.DirEntry) bool {
	return !d.IsDir()
}

func main() {
	home := os.Getenv("HOME")

	dir, err := os.Open(home)
	if err != nil {
		log.Fatal(err)
	}

	fs, err := dir.ReadDir(-1)
	if err != nil {
		log.Fatal(err)
	}

	for _, dirEntry := range fs {
		if IsFile(dirEntry) {
			fmt.Println("file: ", dirEntry.Name())
		} else {
			fmt.Println(" dir: ", dirEntry.Name())
		}
	}
}

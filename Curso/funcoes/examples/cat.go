package main

import (
	"io"
	"log"
	"os"
)

func main() {
	args := os.Args[:]
	if len(args) < 2 {
		log.Fatal("no file specified")
	}
	f, err := os.Open(args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	data := make([]byte, 2048)
	for {
		count, err := f.Read(data)
		os.Stdout.Write(data[:count])
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
	}
}

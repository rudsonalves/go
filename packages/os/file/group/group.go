package main

import (
	"log"
	"os"
)

func main() {
	etcFile := "/etc/group"

	etc, err := os.Open(etcFile)
	if err != nil {
		log.Fatal(err)
	}
	defer etc.Close()

	etcInfo, _ := etc.Stat()
	etcRead := make([]byte, etcInfo.Size())
	n, err := etc.Read(etcRead)????
}

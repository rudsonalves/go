package main

import (
	"log"
	"os"
)

func main() {
	text := []string{
		"While traveling, Einstein wrote daily to his wife Elsa and ",
		"adopted stepdaughters Margot and Ilse. The letters were ",
		"included in the papers bequeathed to the Hebrew University ",
		"of Jerusalem...",
	}

	file, err := os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for _, line := range text {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
}

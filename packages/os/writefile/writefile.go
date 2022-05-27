package main

import (
	"log"
	"os"
)

func main() {
	text := []string{
		"While traveling, Einstein wrote daily to his wife Elsa and",
		"adopted stepdaughters Margot and Ilse. The letters were ",
		"included in the papers bequeathed to the Hebrew University ",
		"of Jerusalem...",
	}

	allText := ""
	for _, line := range text {
		allText += line + "\n"
	}

	data := []byte(allText)
	err := os.WriteFile("test.txt", data, 0600)
	if err != nil {
		log.Fatal(err)
	}
}

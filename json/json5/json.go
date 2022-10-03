package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	fjs, err := os.ReadFile("sample-4.json")
	if err != nil {
		log.Fatal(err)
	}

	var reg map[string]interface{}
	json.Unmarshal(fjs, &reg)

	fmt.Printf("Register:\n Name: %s %s\n Is alive: %v\n Age: %v\n Spouse: %v\n", reg["firstName"], reg["lastName"], reg["isAlive"], reg["age"], reg["spouse"])
}

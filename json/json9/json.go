package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	fjson, err := os.ReadFile("sample.json")
	if err != nil {
		log.Fatal(err)
	}

	var r map[string]interface{}
	json.Unmarshal(fjson, &r)

	fmt.Printf("Name: %s %s\n", r["firstName"], r["lastName"])
	fmt.Printf("Age: %0.f\n\n", r["age"])

	address := r["address"].(map[string]interface{})

	streetAddress := address["streetAddress"]
	city := address["city"]
	state := address["state"]
	postalcode := address["postalCode"]
	fmt.Printf("Address:\n Street: %s\n City: %s\n State: %s\n Postal: %s\n",
		streetAddress, city, state, postalcode)

	children := r["children"].([]interface{})
	fmt.Print("\nChildrens:\n")
	for _, cName := range children {
		fmt.Printf(" %s\n", cName)
	}
}

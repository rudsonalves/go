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

	Children := reg["children"].([]interface{})
	fmt.Println("\nChilden:")
	for _, c := range Children {
		fmt.Printf(" %s\n", c)
	}

	Contacts := reg["phoneNumbers"].([]interface{})
	fmt.Println("\nContacts:")
	for _, c := range Contacts {
		phone := c.(map[string]interface{})
		fmt.Printf(" Type:   %s\n", phone["type"])
		fmt.Printf(" Number: %s\n", phone["number"])
	}

	Address := reg["address"].(map[string]interface{})
	fmt.Println("\nAddress:")
	fmt.Printf(" Street: %s\n City(State): %s (%s)\n Postal code: %s\n", Address["streetAddress"], Address["city"], Address["state"], Address["postalCode"])
}

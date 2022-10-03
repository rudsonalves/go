package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type PhoneItem struct {
	Place  string `json:"type"`
	Number string `json:"number"`
}

type Register struct {
	Name     string      `json:"firstName"`
	LastName string      `json:"lastName"`
	Alive    bool        `json:"isAlive"`
	Age      int         `json:"age"`
	Spouse   string      `json:"spouse"`
	Children []string    `json:"children"`
	Contacts []PhoneItem `json:"phoneNumbers"`
}

func main() {
	fjs, err := os.ReadFile("sample-3.json")
	if err != nil {
		log.Fatal(err)
	}

	reg := Register{}
	json.Unmarshal(fjs, &reg)

	fmt.Printf("Register:\n Name: %s %s\n Is alive: %v\n Age: %d\n Spouse: %v\n", reg.Name, reg.LastName, reg.Alive, reg.Age, reg.Spouse)

	fmt.Println("\nChilden:")
	for _, c := range reg.Children {
		fmt.Printf(" %s\n", c)
	}

	fmt.Println("\nContacts:")
	for _, p := range reg.Contacts {
		fmt.Printf(" Type:   %s\n", p.Place)
		fmt.Printf(" Number: %s\n", p.Number)

	}
}

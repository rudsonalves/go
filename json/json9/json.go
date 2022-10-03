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
	Address  struct {
		Street string `json:"streetAddress"`
		City   string
		State  string
		Postal string `json:"postalCode"`
	}
}

func main() {
	fjs, err := os.ReadFile("sample-4.json")
	if err != nil {
		log.Fatal(err)
	}

	reg := Register{}
	err = json.Unmarshal(fjs, &reg)
	if err != nil {
		log.Fatal(err)
	}

	jsonIndent, err := json.MarshalIndent(reg, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonIndent))

	jsonMarshal, err := json.Marshal(reg)
	if err != nil {
		log.Fatal(err)
	}

	// err = os.WriteFile("test.json", jsonIndent, 700)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	fmt.Println(string(jsonIndent))
	fmt.Println("----------------------")
	fmt.Println(string(jsonMarshal))
}

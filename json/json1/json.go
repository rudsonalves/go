package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Register struct {
	Name     string `json:"firstName"`
	LastName string `json:"lastName"`
	Alive    bool   `json:"isAlive"`
	Age      int    `json:"age"`
}

func main() {
	fjs, err := os.ReadFile("sample-0.json")
	if err != nil {
		log.Fatal(err)
	}

	reg := Register{}
	json.Unmarshal(fjs, &reg)

	fmt.Printf("Register:\n Name: %s %s\n Is alive: %v\n Age: %d\n", reg.Name, reg.LastName, reg.Alive, reg.Age)
}

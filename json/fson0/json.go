package main

import (
	"encoding/json"
	"fmt"
)

type Register struct {
	FirstName string
	LastName  string
	IsAlive   bool
	Age       int
}

func main() {
	// fjs, err := os.ReadFile("sample-0.json")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	JSon := `{"firstName": "John","lastName": "Smith","isAlive": true, "age": 27}`
	reg := Register{}
	json.Unmarshal([]byte(JSon), &reg)

	fmt.Printf("Register:\n Name: %s %s\n Is alive: %v\n Age: %d\n", reg.FirstName, reg.LastName, reg.IsAlive, reg.Age)
}

package main

import "fmt"

// If you know how many key-value pairs you intend to put in the map,
// but don't know the exact values, you ca use make to create a map with
// a default size.
func main() {
	ages := make(map[string][]string, 10)

	ages["dad"] = []string{"Rudson", "Alves"}
	ages["wife"] = []string{"Juliana", "Boasquevisque"}
	ages["son"] = []string{"Gabriel", "Bortolini", "Alves"}

	fmt.Println(ages)
}

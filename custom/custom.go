// +build pro

package main

import "fmt"

var features = []string{
	"Free Feature #1",
	"Free Feature #2",
}

func init() {
	features = append(features,
		"Pro Features #1",
		"Pro Features #2",
	)
}

func main() {
	for _, f := range features {
		fmt.Println(">", f)
	}
}

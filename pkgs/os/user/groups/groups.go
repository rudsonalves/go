package main

import (
	"fmt"
	"log"
	"os/user"
)

func main() {
	g, err := user.LookupGroup("rudson")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("id: %T > %v", g.Gid, g.Gid)
}

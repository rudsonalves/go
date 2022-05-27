package main

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"time"
)

func main() {
	cnt, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	err := exec.CommandContext(cnt, "sleep", "5").Run()
	if err != nil {
		fmt.Println("Fail after 100 milliseconds!!!")
		log.Fatal(err)
	}

	fmt.Println("Sleep 5 sec finish...")
}

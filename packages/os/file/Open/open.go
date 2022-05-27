package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("/etc/group")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	fmt.Println(file.Name())
	fileInfo, _ := file.Stat()
	fmt.Println("   File Name:", fileInfo.Name())
	fmt.Println("        Size:", fileInfo.Size())
	fmt.Println("      Is dir:", fileInfo.IsDir())
	fmt.Println("        Mode:", fileInfo.Mode())
	fmt.Println("         Sys:", fileInfo.Sys())
	fmt.Println(" Modify time:", fileInfo.ModTime())
}

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	home := os.Getenv("HOME")
	path := os.Getenv("PATH")
	gopath := os.Getenv("GOPATH")
	gobin := os.Getenv("GOBIN")

	fmt.Printf("     HOME: %s\n", home)

	pathList := strings.Split(path, ":")
	lenPathList := len(pathList)
	fmt.Printf("PATH[-2:]: %v\n", pathList[lenPathList-2:])

	fmt.Printf("   GOPATH: %s\n", gopath)
	fmt.Printf("    GOBIN: %s\n", gobin)
}

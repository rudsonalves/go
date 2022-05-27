package main

import (
	"fmt"
	"os"
)

func main() {
	// fmt.Println("Page Size:", os.Getpagesize())
	// list, _ := os.Getgroups()
	// fmt.Println("Groups list:", list)
	fmt.Println("           Group id:", os.Getgid())
	// fmt.Println(" Executing Group id:", os.Getegid())
	fmt.Println("          Group pid:", os.Getpid())
	fmt.Println("         Group ppid:", os.Getppid())
	fmt.Println("          Group uid:", os.Getuid())
	groupsList, _ := os.Getgroups()
	fmt.Println("        Groups list:", groupsList)
	dirString, _ := os.Getwd()
	fmt.Println("                pwd:", dirString)
	hostname, _ := os.Hostname()
	fmt.Println("           hostname:", hostname)
}

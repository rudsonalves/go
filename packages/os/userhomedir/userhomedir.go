package main

import (
	"fmt"
	"os"
)

func main() {
	home, _ := os.UserHomeDir()
	conf, _ := os.UserConfigDir()
	cache, _ := os.UserCacheDir()
	fmt.Println(home)
	fmt.Println(conf)
	fmt.Println(cache)
}

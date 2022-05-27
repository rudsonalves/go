package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("ls", "/root/", "/usr/")

	// sdtout, _ := cmd.Output()
	sdtout, _ := cmd.CombinedOutput()

	fmt.Printf("%s\n", sdtout)
}

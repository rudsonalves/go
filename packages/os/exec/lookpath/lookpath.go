package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmds := []string{"ls", "code", "guru", "g++", "blender"}

	for _, cmd := range cmds {
		path, err := exec.LookPath(cmd)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Command %q find in %q\n", cmd, path)
		}
	}
}

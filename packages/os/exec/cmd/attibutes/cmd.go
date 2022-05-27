package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("head", "-n", "5", "/proc/cpuinfo")

	out := bytes.Buffer{}
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("cmd.Args: %v\n", cmd.Args)
	fmt.Printf("cmd.Dir: %v\n", cmd.Dir)
	fmt.Printf("cmd.Env: %v\n", cmd.Env)
	fmt.Printf("cmd.ExtraFiles: %v\n", cmd.ExtraFiles)
	fmt.Printf("cmd.Path: %v\n", cmd.Path)
	fmt.Printf("cmd.Process: %v\n", cmd.Process)
	fmt.Printf("cmd.Process Pid: %v\n", cmd.ProcessState.Pid())
	fmt.Printf("cmd.Process Success: %v\n", cmd.ProcessState.Success())
	fmt.Printf("\ncmd.Stdout:\n%v\n", cmd.Stdout)
}

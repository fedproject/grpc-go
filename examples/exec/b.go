package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// come out of package b and then go inside package a to run the executable file as
	cmd := exec.Command("go ", "run ", "aaa/aaa.go")
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	}
}

package main

import (
	"fmt"
	"os/exec"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	cmd := exec.Command("./../other/other", "hello", "world")
	data, err := cmd.Output()
	checkError(err)
	fmt.Println(string(data))
}

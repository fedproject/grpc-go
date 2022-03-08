package main

import (
	"fmt"
	"os/exec"
	"../route_guide/client"
)

func main() {
	// cmnd := exec.Command("go", "run", "o_flag.go", "-age 1")
	// //cmnd.Run() // and wait
	// cmnd.Start()
	client.main_c("N01")

	cmnd := exec.Command("go run o_flag.go", " -age 1")
	// if err == nil {
	// 	log.Println(cmnd)
	// }
	if err := cmnd.Run(); err != nil {
		fmt.Println(err)
	}
	// log.Println(err)
	// log.Println("log")
}

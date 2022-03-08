package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("命令行的参数数目：", len(os.Args))
	for i, v := range os.Args {
		if i == 0 {
			continue
		}
		fmt.Printf("args[%v]=%v\n", i, v)
	}
}

// flags.exe -h
package main

import (
	"flag"
	"fmt"
)

func main() {
	namePtr := flag.String("name", "AR", "name")
	agePtr := flag.Int("age", 3700, "age")
	flag.Parse()
	fmt.Println(*namePtr, *agePtr) //AR 3700
}

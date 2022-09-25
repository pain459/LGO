package main

import (
	"fmt"
)

func main() {
	fmt.Println("Program starts here.")
	x := 0
	// y := 0
	for {
		if x++; x > 2 {
			fmt.Println(x)
			break
		}
	}
}

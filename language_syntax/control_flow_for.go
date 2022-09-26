package main

import (
	"fmt"
)

func main() {
	fmt.Println("Program starts here.")
	x := 0
	y := 0
	for {
		if x++; x > 10 {
			fmt.Println(x)
			break
		}
	}

	for y < 3 {
		fmt.Println(y)
		y++
	}

	for z := 0; z < 10; z++ {
		if z < 8 {
			continue // give back the control to for loop.
		}
		fmt.Println(z)
	}

	fmt.Println("New test loop starts here.")
	for a := 0; a < 15; a++ {
		if a < 10 {
			fmt.Println(a)
		}
	}
}

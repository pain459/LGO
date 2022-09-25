package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Program starts here.\n")
	var isBool bool // default to false
	var isBig bool = true
	var isBlue, isRed bool = true, false
	fmt.Println(isBool)
	fmt.Println(isBig)
	if isBig {
		fmt.Println("Printing as isBig is True now.")
	}
	if !isBool {
		fmt.Println("Printing as isBook is now made as True.")
	}
	fmt.Println(isBool || isBig)
	fmt.Println(isBig && isBool)
	fmt.Println(!isBig)
	fmt.Println(isBool || isBig && !isBool)
	fmt.Println(isBlue || isRed)
}

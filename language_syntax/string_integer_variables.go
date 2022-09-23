package main

import (
	"fmt"
)

var Num1 = 199

func main() {
	Num2 := 299
	fmt.Println("This is a base program")
	fmt.Printf("This is Num1 %v\n", Num1)
	fmt.Printf("%v\n", Num2)
	Num1 = Num1 * Num2
	fmt.Printf("This is the new result %v\n", Num1)
}

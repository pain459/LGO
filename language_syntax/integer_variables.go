package main

import (
	"fmt"
)

var num1 int
var num2 int = 100

func main() {
	fmt.Println("Program starts now.")
	fmt.Println(num1)
	fmt.Println(num2)
	num3 := 200
	fmt.Println(num3)
	num3 += num2
	fmt.Println(num3)
	num3++
	fmt.Println(num3)

}

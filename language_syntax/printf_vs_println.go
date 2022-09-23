package main

import "fmt"

var String1 string = "Sammy"
var Num1 int = 299

func main() {
	String2 := "Maddy"
	fmt.Println("Program starts here.")
	fmt.Printf("This is the global string %v\n", String1)
	fmt.Println("This is variable in func", String2)
	Num1++
	fmt.Println("Incremented number is", Num1)
	Num1--
	fmt.Printf("Earlier number is %v\n", Num1)
}

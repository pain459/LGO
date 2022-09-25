package main

import (
	"fmt"
)

func main() {
	fmt.Println("Program starts here.")
	var population int = 50000
	if population <= 500 {
		fmt.Println("Population is", population, "in the allowed range.")
	} else if population == 5000 {
		fmt.Println("Population is exactly", population)
	} else {
		fmt.Println("No census data.")
	}
	if toggle := true; toggle {
		fmt.Println("Toggle is declared as", toggle)
	}
}

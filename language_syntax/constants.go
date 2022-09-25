package main

import (
	"fmt"
)

const yearNum int = 2022

func main() {
	fmt.Println("Program starts here.")
	if yearNum == 2022 {
		fmt.Println("The year is 2022")
		const leapYear = "nope"
		if leapYear == "nope" {
			fmt.Println("2022 is a", leapYear)
		}
	}
	fmt.Println(yearNum)
	// fmt.Println(leapYear)  // Constant out of scope and hence the error.
}

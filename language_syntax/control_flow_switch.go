package main

import (
	"fmt"
)

func main() {
	fmt.Println("Program execution starts now.")
	switch signal := 1; signal {
	case 0:
		fmt.Println("Case 0 executed.")
	case 1:
		fmt.Println("Case 1 executed.")
	case 2:
		fmt.Println("Case 2 executed.")
	}
	var score = 25

	switch {
	case score > 100:
		fmt.Println("Score is greater than", score)
	case score >= 50:
		fmt.Println("Score is grater than", score)
	case score < 50:
		fmt.Println("Score is not greater than", score)
	}
}

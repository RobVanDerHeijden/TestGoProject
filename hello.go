package main

// Importing Packaging (fmt is a basic package for I/O text)
import (
	"fmt"
)

func main() {
	// Basic Print function
	fmt.Println("Hello, World3!")

	// Variables with short tag syntax   :=
	x := 6
	y := 1
	sum := x + y
	sumEvaluationText := ""
	sumTextOutput := ""

	// If statement + if else + else
	if sum > 8 {
		sumEvaluationText = "is higher then 8"
	} else if sum < 8 {
		sumEvaluationText = "is lower then 8"
	} else {
		sumEvaluationText = "is equal to 8"
	}

	// Sprintf for string/int concatenation (%d = int, %s = string)
	sumTextOutput = fmt.Sprintf("%d: %s", sum, sumEvaluationText)
	fmt.Println(sumTextOutput)

	//fmt.Println(sum)

}

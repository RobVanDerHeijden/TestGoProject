package main

// Importing Packaging (fmt is a basic package for I/O text)
import (
	"errors"
	"fmt"
	"math"
)

// Struct (kinda object without methods)
type person struct {
	name string
	age  int
}

func main() {
	// Basic Print function
	fmt.Println("Hello, World!")

	// Making struct variable
	p := person{name: "Rob", age: 28}
	fmt.Println(p.age, p.name)

	// Pointer simple example (& before variable gives memory example)
	pointerExample := 9
	incCalc(&pointerExample)
	fmt.Println(pointerExample)

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

	// Array (fixed lenght)
	arrayExample := [5]int{3, 9, 46, 842, 31}
	fmt.Println(arrayExample)

	// Slice (no fixxed lenght)
	sliceExample := []int{3, 9, 46, 842, 31}
	sliceExample = append(sliceExample, 357)
	fmt.Println(sliceExample)

	// Map -> map[keys]values
	mapExample := make(map[string]int)
	mapExample["apples"] = 13
	mapExample["bananas"] = 5
	mapExample["pineapples"] = 2
	fmt.Println(mapExample)
	fmt.Println(mapExample["apples"])
	delete(mapExample, "bananas")
	fmt.Println(mapExample)

	// Loop (for loop)
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	// Loop (while loop)
	i := 0
	for i < 3 {
		fmt.Println(i)
		i++
	}
	// Loop (Array: iterate over each element)
	arrayIterateExample := [5]string{"yes", "no", "maybe", "perhps", "of course"}
	for index, value := range arrayIterateExample {
		fmt.Println("index", index, "value", value)
	}
	// Loop (Map: iterate over each element)
	mapIterateExample := make(map[string]string)
	mapIterateExample["chocolate"] = "Delicious"
	mapIterateExample["strawbarries"] = "also good"
	for key, value := range mapIterateExample {
		fmt.Println("key", key, "value", value)
	}

	// Call outside basic additionfunction
	resultSumCalc := sumCalc(314, 159)
	fmt.Println(resultSumCalc)

	// Call outside square root function
	resultsqrtCalc, err := sqrtCalc(25)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultsqrtCalc)
	}

}

// Seperate function basic addition
func sumCalc(x int, y int) int {
	return x + y
}

func sqrtCalc(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undifined for negative numbers")
	}
	return math.Sqrt(x), nil
}

// Used for pointer example (* is for dereferencing)
func incCalc(x *int) {
	*x++
}

package main

import (
	"fmt"
)

// func return multiple value
// define same multiple parameters data type at the end
func swap(a, b int) (int, int) {
	return b, a
}

// func return single value
func sum(a float64, b float64) float64 {
	return a + b
}

// func return single value
func max(num1, num2 int) int {
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}

	return result
}

// func return predefined value
func square(a int) (area int) {
	area = (a * a)
	return
}

// variadic parameters (unlimited parameter values)
func sumNumbers(numbers ...int) (result int) {
	fmt.Println(numbers)
	for _, number := range numbers {
		result += number
	}
	return result
}

// closure function, function that assigned to a variable
var squareArea = func(a int) (area int) {
	area = a * a
	return
}

// immediately invoked, below variable "greet" will have value "Hello John"
var greet = func(name string) string {
	return "Hello " + name
}("John")

func main() {
	fmt.Println(sum(0.5, 0.1))
	max(1, 2)
	swap(1, 2)
	square(5)
	sumNumbers(1, 2, 3)
	numbers := []int{1, 2, 3}
	sumNumbers(numbers...)
	squareArea(5)
	println(greet)
}

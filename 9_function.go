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

// func as parameter
func greeting_student(name string, callback func(string) bool) string {
	var result string
	if greeting_type := callback("student"); greeting_type {
		result = "Hello Student " + name
		return result
	}
	result = "Hello People " + name
	return result
}

func main() {
	// func return single value
	max(1, 2)

	// func return multiple value
	swapResult1, swapResult2 := swap(1, 2)
	println(swapResult1, swapResult2)

	// func return predefined value
	squareResult := square(5)
	println(squareResult)

	// variadic parameters (unlimited parameter values)
	sumNumbersResult := sumNumbers(1, 2, 3)
	numbers := []int{1, 2, 3}
	sumNumbersResult2 := sumNumbers(numbers...)
	println(sumNumbersResult, sumNumbersResult2)

	// closure function, function that assigned to a variable
	squareAreaResult := squareArea(5)
	println(squareAreaResult)

	// immediately invoked, below variable "greet" will have value "Hello John"
	println(greet)

	// func as parameter
	greeting_student_result := greeting_student("John", func(s string) bool {
		if s == "student" {
			return true
		} else {
			return false
		}
	})
	println(greeting_student_result)
}

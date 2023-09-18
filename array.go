package main

import (
	"fmt"
	"reflect"
)

func main() {
	// create variable array
	var scores [5]int
	fmt.Println("should return zero valued arrays", scores)

	// assign value to an element of array by index
	scores[0] = 100
	scores[1] = 99
	fmt.Println(scores)

	// built in function len() to know length number of an array
	var count_scores = len(scores)
	fmt.Println("should return count value of scores", count_scores)

	// declaration variable with array int
	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println(reflect.TypeOf(numbers))
	fmt.Println(numbers)

	// create multidimensional array
	var array2D [1][2]int
	fmt.Println(array2D)
	array2D[0][0] = 1
	array2D[0][1] = 2
	fmt.Println(array2D)

	var array2D_2 = [1][2]int{
		{
			1, 2,
		},
	}
	fmt.Println(array2D_2)

}

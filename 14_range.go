package main

import "fmt"

/*
Range
1. iterate over items of array, slice, map, channel
2. array & slice return index as integer and value
3. map return key and value

*/
func main() {

	// slice
	numbers := []int{10, 20, 30, 40}

	for index := range numbers {
		fmt.Println(index)
	}

	for index, value := range numbers {
		fmt.Println(index, value)
	}

	// map
	student := make(map[string]string, 3)
	student["name"] = "abc"
	student["address"] = "def"

	for k, v := range student {
		fmt.Println(k, v)
	}

}

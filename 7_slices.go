package main

import (
	"fmt"
	"strings"
)

func main() {
	var line = (strings.Repeat("=", 100) + "\n")
	// slice is reference of array
	// to create a slice no need to define the length of slice

	// create array
	fmt.Println("create array")
	var fruits_array = [5]string{"apple", "pineapple", "watermelon", "banana", "grape"}
	fmt.Println("fruits_array", fruits_array)
	fmt.Println(line)

	// create slice
	fmt.Println("create slice")
	var fruits_slice = []string{"guava", "cherry", "mango", "durian", "jackfruit"}
	fmt.Println("fruits_slice", fruits_slice)
	fmt.Println(line)

	// create slice reference to an array
	fmt.Println("create slice reference to an array")
	var fruits_slice_reference_array = fruits_array[0:3]
	fmt.Println("fruits_slice_reference_array", fruits_slice_reference_array)
	fruits_slice_reference_array[0] = "orange"
	fmt.Println("fruits_slice_reference_array after change index 0", fruits_slice_reference_array)
	fmt.Println("fruits_array after change fruits_slice_reference_array index 0 ", fruits_array)
	fmt.Println(line)

	// make() create slice without array
	fmt.Println("make() create slice without array")
	fmt.Println("make(type, len, cap)")
	days := make([]string, 2, 3)
	// days[0] = "Monday"
	// days[1] = "Tuesday"
	fmt.Println("days", days)
	fmt.Println("len", len(days))
	fmt.Println("cap", cap(days))
	var append_days_1 = append(days, "Friday")
	var append_days_2 = append(days, "Sunday")
	fmt.Println("append_days_1", append_days_1)
	fmt.Println("append_days_2", append_days_2)
	fmt.Println("days", days)
	fmt.Println(line)

	// len() count total element of a slice
	fmt.Println(len(fruits_slice))
	fmt.Println(len(fruits_slice_reference_array))

	// cap() to know max capacity of a slice
	fmt.Println(cap(fruits_slice))
	fmt.Println(cap(fruits_slice_reference_array))

	// append() a slice
	list_fruit := []string{"orange", "grape", "mango"}
	fmt.Println("list_fruit", list_fruit)
	fmt.Println("cap", cap(list_fruit))
	fmt.Println("len", len(list_fruit))
	list_fruit_a := list_fruit[1:3] // change to [1:2] to show reference between slice
	fmt.Println("list_fruit_a", list_fruit_a)
	fmt.Println("cap", cap(list_fruit_a))
	fmt.Println("len", len(list_fruit_a))
	list_fruit_b := append(list_fruit_a, "papaya")
	fmt.Println("list_fruit_b", list_fruit_b)
	fmt.Println("cap", cap(list_fruit_b))
	fmt.Println("len", len(list_fruit_b))
	list_fruit_b[0] = "xxx"
	fmt.Println(list_fruit_a)
	fmt.Println(list_fruit_b)

	// copy()
	source := []string{"apple", "orange", "guava"}
	destination := make([]string, 1)
	n := copy(destination, source)
	fmt.Println("source", source)
	fmt.Println("destination", destination)
	fmt.Println(n)

	// access slice with 3 index - 3 index adalah teknik slicing elemen yang sekaligus menentukan kapasitasnya
	var fruits = []string{"apple", "grape", "banana"}
	var aFruits = fruits[0:2]
	var bFruits = fruits[0:2:2]

	fmt.Println(fruits)      // ["apple", "grape", "banana"]
	fmt.Println(len(fruits)) // len: 3
	fmt.Println(cap(fruits)) // cap: 3

	fmt.Println(aFruits)      // ["apple", "grape"]
	fmt.Println(len(aFruits)) // len: 2
	fmt.Println(cap(aFruits)) // cap: 3

	fmt.Println(bFruits)      // ["apple", "grape"]
	fmt.Println(len(bFruits)) // len: 2
	fmt.Println(cap(bFruits)) // cap: 2

}

package main

import "fmt"

/*
Map is data type key-value pair.
*/

func main() {
	// create map
	employee := map[string]string{
		"name":    "Taylor",
		"address": "New Jersey",
	}
	fmt.Println(employee)
	fmt.Println("***** Add key-value to a map object")
	employee["number"] = "1234"
	fmt.Println(employee)
	fmt.Println(len(employee)) // total key-value

	// checking a key exist or not in a map
	student := make(map[string]string)
	student["fullname"] = "August"
	value, ok := student["fullname"]
	if ok {
		fmt.Println("fullname", value)
	} else {
		fmt.Println("fullname is not present")
	}

	// delete function
	delete(student, "fullname")
	value, ok = student["fullname"]
	if ok {
		fmt.Println("fullname", value)
	} else {
		fmt.Println("fullname is not present")
	}

}

package main

import "fmt"

func main() {
	fmt.Println("Learn Data Type: Map")

	employee := map[any]string{
		"name":    "Taylor",
		"address": "New Jersey",
	}

	fmt.Println(employee)

	employee["number"] = "1234"

	fmt.Println(employee)
	fmt.Println(employee["name"])
	fmt.Println(len(employee)) // total key-value

	students := make(map[string]string)
	students["fullname"] = "August"

	value, bool := students["fullname"]

	fmt.Println(value, bool)

	fmt.Println(students)
	fmt.Println(students["fullname"])
}

package main

import (
	"fmt"
	"math"
)

func main() {
	// var variable_list optional_data_type;
	// optional_data_type can be byte, int, float32, complex64, boolean or any user-defined object.

	// variable definition
	var age int
	var full_name string
	var status_happy bool
	var salary float32
	fmt.Println(age, full_name, status_happy, salary)

	// dynamic type declaration
	var postal_code = 59111
	var street = "Jalan Kenangan"
	fmt.Println(postal_code, street)

	// static type declaration
	var surname string
	surname = "nic"
	fmt.Println(surname)

	// dynamic type declaration / type inference
	var x float64 = 20.0
	y := 20.0
	fmt.Println(x, y)
	fmt.Printf("x is of type %T \n", x) // y is of type float64
	fmt.Printf("y is of type %T \n", y) // x is of type float64

	// multiple declaration variables
	var postcode = 59111
	var (
		city    = "Pati"
		country = "Indonesia"
	)

	fmt.Println(city, country, postcode)

	// constant
	const name string = "keni"
	fmt.Println(name)
	// name := "test" --> this will return error cannot assign to name
	const number = 50
	sin := math.Sin(number)
	fmt.Printf("%v", sin)

	// multiple declaration constant
	const (
		employee bool = true
		id_card  int  = 30
	)

	fmt.Println(employee)
	fmt.Println(id_card)
}

package main

import "fmt"

func main() {

	fmt.Println("*********Example 1")
	i, j := 10, 789

	p := &i         // var p has value address of var i
	fmt.Println(p)  // output address
	fmt.Println(j)  // output value
	fmt.Println(*p) // output value of i trough p
	*p = 20         // dereferencing or indirecting
	fmt.Println(i)  // output 20
	fmt.Println(p)  // output address of i
	fmt.Println(*p) // output value of i 20

	fmt.Println("*********Example 2")
	myInitialValue := 20         // type int
	myPointer := &myInitialValue // type *int
	fmt.Println(myPointer)       // should return address of myInitialValue
	fmt.Println(*myPointer)      // should return the value of myInitialValue variable's address

	fmt.Println("*********Example 3")
	numbers := []int{1, 2, 3, 4, 5} // create slice of int
	fmt.Println(numbers)
	newNumbers := numbers // create a var newNumbers with numbers as value and it will be a pointer to numbers
	fmt.Println(newNumbers)
	newNumbers[0] = 999 // as a pointer of numbers it will also change the first index of numbers
	fmt.Println(newNumbers)
	fmt.Println(numbers)

	fmt.Println("*********Example 4")
	type Student struct {
		value int
	}
	str := new(Student)
	var clonedStr = str
	clonedStr.value = 60
	fmt.Println(str)

}

package main

import "fmt"

func main() {
	/* Learn golang */
	fmt.Println("hello world")
	fmt.Println("hello " + "world")
	fmt.Println("hello", "world")

	fmt.Println("2+2=", 2+2)
	fmt.Println("2/2=", 2/2)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
	fmt.Println(!false)

	// zero value
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

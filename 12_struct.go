package main

import (
	"fmt"
	"time"
)

// struct collection of fields
type Book struct {
	title  string
	id     string
	author string
	year   int
}

func main() {

	fmt.Println("*********Example 1")
	world := Book{"World", "2", "noname", 2000}
	fmt.Println(world)

	fmt.Println("*********Example 2")
	var myBook Book
	myBook.author = "noname"
	myBook.id = "1"
	myBook.title = "My Book"
	myBook.year = 1990
	fmt.Println(myBook)

	fmt.Println("*********Example 3")
	fmt.Println("*********Pointer Struct")
	fmt.Println("Create a struct")
	globe := Book{
		title:  "Globe",
		id:     "3",
		author: "noname",
		year:   1990,
	}
	fmt.Println("Create a pointer")
	pointerGlobe := &globe
	fmt.Println(pointerGlobe)
	pointerGlobe.author = "BUDI"
	fmt.Println(pointerGlobe.author)
	fmt.Println(pointerGlobe)
	fmt.Println(globe)
	fmt.Println("Dereference pointer")
	derefGlobe := *pointerGlobe
	derefGlobe.author = "ANAK BUDI"
	fmt.Println(derefGlobe)
	fmt.Println(pointerGlobe)
	fmt.Println(globe)

	fmt.Println("*********Example 4")
	fmt.Println("*********Struct Literals") // create struct with single syntax
	var water Book                          // declaring a variable with type Book and values set to zero values golang
	fmt.Println(water)
	var bottle Book = Book{"Bottle", "4", "noname", 2009} // create instance of a struct type using a struct literal
	fmt.Println(bottle)

	fmt.Println("*********Example 5")
	items := []struct {
		title   string
		expired time.Time
	}{
		{title: "abc", expired: time.Now()}, {title: "def", expired: time.Now()},
	}
	fmt.Println(items)
}

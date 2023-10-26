package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	var text = "Final Score"
	var score = 101
	if score <= 60 {
		fmt.Println(text, "D")
	} else if score <= 80 {
		fmt.Println(text, "C")
	} else if score <= 90 {
		fmt.Println(text, "B")
	} else if score <= 100 {
		fmt.Println(text, "A")
	} else {
		fmt.Println("Out of scope")
	}

	// if with shorthand
	type Book struct {
		Title string
	}
	obj := Book{}
	if err := json.Unmarshal([]byte("abc"), obj); err != nil {
		fmt.Println(err)
	}
	fmt.Println(obj)

	// switch case
	switch true {
	case score <= 60:
		fmt.Println(text, "D")
	case score <= 80:
		fmt.Println(text, "C")
	case score <= 90:
		fmt.Println(text, "B")
	case score <= 100:
		fmt.Println(text, "A")
	default:
		fmt.Println("Out of scope")
	}

}

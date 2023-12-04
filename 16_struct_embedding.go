package main

import (
	"fmt"
)

type base struct {
	name string
}

func (b base) describe() string {
	return b.name
}

type container struct {
	base
	str string
}

func example1() {
	fmt.Println("> Example 1")

	container1 := container{
		base: base{
			name: "base1",
		},
		str: "container1",
	}

	fmt.Println(container1.name)
	fmt.Println(container1.describe())
	fmt.Printf("%+v\n", container1)
}

type Author struct {
	AuthorName string
}

type Publisher struct {
	Publisher string
	Address   string
}

type Book struct {
	Title string
	Year  int
	Author
}

type AuthorBooks struct {
	Author
	Books []Book
}

func (a Author) Greet() string {
	return fmt.Sprintf("Hi, I'm %s", a.AuthorName)
}

func (b Book) GetTitle() string {
	return b.Title
}

func example2() {
	fmt.Println("> Example 2")
	book1 := Book{
		Title: "abcd",
		Year:  2090,
		Author: Author{
			AuthorName: "john",
		},
	}

	fmt.Printf("%+v\n", book1)
	fmt.Printf("The result would be equal '%s' and '%s'\n", book1.Author.Greet(), book1.Greet())
}

type BookInformation interface {
	BookDetail() string
}

func (b Book) BookDetail() string {
	return fmt.Sprintf("Title book %s and the author is %s", b.Title, b.AuthorName)
}

func example3() {
	fmt.Println("> Example 3")
	book1 := Book{
		Title: "abcd",
		Year:  2090,
		Author: Author{
			AuthorName: "john",
		},
	}

	bookDetail := BookInformation.BookDetail(book1)
	fmt.Println(bookDetail)
}

func main() {
	fmt.Println("Struct Embedding")
	example1()
	example2()
	example3()
}

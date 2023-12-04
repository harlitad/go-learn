package main

import (
	"fmt"
)

type geometry interface {
	area() float64
}

type circle struct {
	width, lenght float64
}

type rectangle struct {
	width, lenght float64
}

func (r rectangle) area() float64 {
	return r.width * r.lenght
}

func (c circle) area() float64 {
	return c.width * c.lenght
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
}

func main() {

	r := rectangle{4, 4}
	c := circle{2, 3}

	measure(r)
	measure(c)
}

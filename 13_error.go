package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	// handle error
	m, err := GreetV2("abc", 0)
	if err != nil {
		log.Fatal(err)
	}

	m, err = Greet("", 10)
	if err != nil {
		// log and terminate
		log.Fatal(err)
	}
	fmt.Println(m)
}

func Greet(name string, age int) (string, error) {
	if name == "" {
		return "", errors.New("name shouldn't be empty")
	}

	if age == 0 {
		return "", fmt.Errorf("age can't be empty, %d", age)
	}
	return fmt.Sprintf("Welcome %s with age %d", name, age), nil
}

// custom error should implement method Error() and return string
type GreetNameError struct{}
type GreetAgeError struct{}

func (g GreetNameError) Error() string {
	return "name should not be empty"
}

func (g GreetAgeError) Error() string {
	return "age should not be empty"
}

func GreetV2(name string, age int) (string, error) {
	if name == "" {
		return "", GreetNameError{}
	}

	if age == 0 {
		return "", GreetAgeError{}
	}

	return fmt.Sprintf("Welcome %s", name), nil
}

package main

import (
	"fmt"
	"time"
)

func Hello(name string) {
	fmt.Println("Hello", name, time.Now())
	time.Sleep(2 * time.Second)
}

func f(from string) {
	for i := 0; i < 3; i++ {
		go Hello(fmt.Sprint(i))
	}
}

func main() {
	example1()
}

func example1() {
	fmt.Println("> Example 1")
	f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(2 * time.Second)
	fmt.Println("done")

	/* expected output:
	going
	Hello 2 2024-01-05 13:23:22.129993811 +0700 +07 m=+0.000218301
	Hello 0 2024-01-05 13:23:22.129890811 +0700 +07 m=+0.000115201
	Hello 1 2024-01-05 13:23:22.129968811 +0700 +07 m=+0.000193301
	done
	*/
}

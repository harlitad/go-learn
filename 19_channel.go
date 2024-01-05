package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	example1()
	example2()
}

func Hello(name string, channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- fmt.Sprintf("Hi %s", name)
}

func Hello2(name string, channel chan<- string, wg *sync.WaitGroup) {
	time.Sleep(2 * time.Second)
	channel <- fmt.Sprintf("Hi %s", name)
	wg.Done()
}

func example1() {
	fmt.Println("> Example 1")
	start := time.Now()
	fmt.Println(">> Start", start)

	// create channel
	channel := make(chan string)
	defer close(channel)

	go Hello("Budi", channel)

	res := <-channel
	fmt.Println(res)

	fmt.Println(">> Finish", time.Since(start)) // running 1 goroutine within 2 seconds duration
}

func example2() {
	fmt.Println("> Example 2")
	start := time.Now()
	fmt.Println(start)

	// create channel
	channel := make(chan string)
	// waitGroup to ensure that all goroutine finished
	wg := &sync.WaitGroup{}

	for _, name := range []string{"Budi", "John"} {
		wg.Add(1)
		go Hello2(name, channel, wg)
	}

	go func() {
		wg.Wait()
		close(channel)
	}() // wait all goroutine and close channel

	for res := range channel {
		fmt.Println(res)
	}

	fmt.Println(">> Finish", time.Since(start)) // running 2 goroutine within 2 seconds duration
}

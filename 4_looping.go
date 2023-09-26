package main

import "fmt"

func looping_for(count int) {
	for i := 0; i < count; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("First Loop:", i)
	}

	number := 0
	for number < count {
		fmt.Println("Second Loop:", number)
		number++
	}

	numbers := [10]int{}

	for index, value := range numbers {
		fmt.Println("Third Loop:", index, value)
	}
}

func infinite_loop() {
	for true {
		fmt.Println("Infinite Loop")
	}
}

func loop_control_statement_break() {
	/* statement break uses to terminate loop process */
	i := 0
	for i < 5 {
		if i == 3 {
			fmt.Printf("%d == 3", i)
			break
		}
		fmt.Println(i)
		i++
	}
}

func loop_control_statement_continue() {
	/* statement continue uses to skip current loop process to the next */
	i := 0
	for i < 5 {
		if i == 3 {
			i++
			continue
		}
		fmt.Println("index is ", i)
		i++
	}
}

func main() {
	looping_for(5)
	// infinite_loop()
	loop_control_statement_break()
	loop_control_statement_continue()

}

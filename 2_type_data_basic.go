package main

import "fmt"

func main() {
	// number
	var number1 = 10
	var number2 = 20
	var sum = number1 + number2
	var number3 = 9.7
	fmt.Println(sum)
	fmt.Println(number3)

	// bool
	var like = true
	var dislike = false
	fmt.Println(like, dislike)

	// string
	var name string = "Harlita"
	fmt.Println(len(name))
	fmt.Println(name[0])

	// type conversion
	var numb32 int32 = 128
	var numb64 int64 = int64(numb32)
	var numb8 int8 = int8(numb64)
	fmt.Println(numb32, numb64, numb8) // belum paham

	var full_name string = "Nicholas"
	var first_char uint8 = full_name[0]
	var second_char byte = full_name[1]
	fmt.Println(first_char)
	fmt.Println(second_char)
	fmt.Println(string(first_char))

}

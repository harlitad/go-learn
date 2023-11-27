package main

import (
	"encoding/json"
	"fmt"
)

type book struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Author string   `json:"author"`
	Tags   []string `json:"tags"`
}

type Student struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

func main() {
	// object to string using marshall
	fmt.Println("************ object to string using marshall")
	fmt.Println("************ example 1")
	rSample1, err := json.Marshal(1)
	fmt.Println(string(rSample1), err)

	fmt.Println("************ example 2")
	rSample2, err := json.Marshal("")
	fmt.Println(string(rSample2), err)

	fmt.Println("************ example 3")
	sample3 := []book{
		{
			Title:  "title 1",
			Author: "author 1",
			Year:   2000,
			Tags:   []string{},
		},
	}
	fmt.Println(sample3)
	rSample3, err := json.Marshal(sample3)
	fmt.Println(string(rSample3), err)

	fmt.Println("************ example 4")
	rSample4, err := json.MarshalIndent(sample3, "", " ")
	fmt.Println(string(rSample4))

	toStringJson := func(obj interface{}) string {
		v, _ := json.MarshalIndent(obj, "", "-")
		return string(v)
	}
	fmt.Println(toStringJson(sample3))

	// json string to object using unmarshall
	fmt.Println("************ json string to object using unmarshall")
	student := Student{}
	strStudent := `{"name": "budi", "id": "1"}`
	// student should be passing as pointer to assign the value in the same address
	if err := json.Unmarshal([]byte(strStudent), &student); err != nil {
		fmt.Println("err: ", err)
	}
	fmt.Println(student.Name, student.Id)

}

package main

import (
	"encoding/json"
	"fmt"
)

// ต้องเป็น pulic ถึงจะเข้าถึงได้ เพราะอยู่คนละ package
type Course struct {
	Name       string
	Level      string
	Views      int
	Instructor string
	FullPrice  int
}

func main() {
	course := Course{
		Name:       "Go",
		Level:      "Beginner",
		Views:      100,
		Instructor: "Doittikorn",
		FullPrice:  1000,
	}

	b, err := json.Marshal(course)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("type : %T \n", b)
	fmt.Printf("byte : %v \n", b)
	fmt.Printf("string : %s \n", b)

}

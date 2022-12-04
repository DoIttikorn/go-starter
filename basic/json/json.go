package main

import (
	"encoding/json"
	"fmt"
)

// ต้องเป็น pulic ถึงจะเข้าถึงได้ เพราะอยู่คนละ package
type Course struct {
	Name       string `json:"name"`
	Level      string `json:"level"`
	Views      int    `json:"views"`
	Instructor string `json:"instructor"`
	FullPrice  int    `json:"full_price"`
}

func main() {
	course := &Course{
		Name:       "Go",
		Level:      "Beginner",
		Views:      100,
		Instructor: "Doittikorn",
		FullPrice:  1000,
	}
	// package json ไม่สามารถเข้าถึง field ที่เป็น private ได้ จึงต้องเปลี่ยนเป็น public
	b, err := json.Marshal(course) // Marshal (การจัดเรียง)  คือการแปลง struct ให้เป็น json
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("type : %T \n", b)
	fmt.Printf("byte : %v \n", b)
	fmt.Printf("string : %s \n", b)

	data := []byte(`{"name":"Go","level":"Beginner","views":100,"instructor":"Doittikorn","full_price":1000}`)

	var c Course
	err = json.Unmarshal(data, &c) // Unmarshal (การแยก) คือการแปลง json ให้เป็น struct

	fmt.Printf("type : %T \n", c)
	fmt.Printf("byte : %v \n", c)
	fmt.Printf("string : %s \n", c)

}

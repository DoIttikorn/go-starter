package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// การประกาศ custom type ใน Go จะเป็นการประกาศ
// struct ใหม่ โดย struct ใหม่นี้จะมี field ที่เป็น type ของ struct ที่เราประกาศไว้ก่อนหน้านี้
type course struct {
	name       string
	instructor string
	price      int
}

func main() {
	// Create a new struct.
	p := Person{
		Name: "John",
		Age:  42,
	}
	// Access the field.
	fmt.Println(p.Name)
	var c = course{name: "Go", instructor: "John", price: 100}
	fmt.Printf("c: %#v\n", c)

	// การเข้าถึง field ของ struct ใน Go จะเป็นการเข้าถึงด้วย . ตามด้วยชื่อ field
	courseName := c.name
	fmt.Println(courseName)
	// การ set ค่าของ struct ใน Go จะเป็นการเข้าถึงด้วย . ตามด้วยชื่อ field
	c.instructor = "John Doe"
	fmt.Printf("c: %#v\n", c)

	var c2 = course{"Kotlin", "Dodo", 200}
	var c3 = course{instructor: "Ploy"}
	var c4 = course{}
	fmt.Printf("c2: %#v\n", c2)
	fmt.Printf("c3: %#v\n", c3)
	fmt.Printf("c4: %#v\n", c4)

}

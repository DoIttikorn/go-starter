package main

import "fmt"

type course struct {
	name, instructor string
	price            int
}

// pass by value c ใน discount จะเป็น copy ของ c ที่อยู่ใน main และเปลี่ยนค่าใน discount ไม่มีผลต่อค่าใน main
// func discount(c course) int {
// 	c.price = c.price - 599
// 	fmt.Println("discount", c.price)
// 	return c.price
// }

func discount(c *course) int {
	// (*c).price = (*c).price - 599 // ถ้าเป็นแบบนี้ go จะเข้าใจว่าเป็นการเรียกใช้ field ของ struct แทน
	c.price = c.price - 599
	fmt.Println("discount", c.price)
	return c.price
	// ค่าที่ return ออกมาจะเป็นค่าใน c.price ที่ถูกลบแล้ว ถ้าเป็น pointer จะเป็นค่าที่อยู่ใน address นั้น
	// แต่ถ้า return ออกมาเป็น struct ที่ไม่ใช่ pointer จะเป็นค่าที่ถูก copy มาจาก struct นั้น
}

func (c *course) discount2(d int) int {
	c.price = c.price - d
	fmt.Println("discount", c.price)
	return c.price
}

func main() {
	// c := new(course) // สร้าง pointer ของ course แล้วเก็บค่าที่อยู่ของ course ที่สร้างใหม่ไว้ใน c
	c := &course{
		name:       "Go",
		instructor: "Mihalis",
		price:      10000,
	}
	fmt.Printf("c: %#v\n", c)

	d := discount(c) // ส่งค่า address ของ c ไปให้ discount
	fmt.Println("discount price from func ", d)
	fmt.Println("struct property price: ", c.price)

	f := c.discount2(120) // method ที่เป็น pointer จะเปลี่ยนค่าใน struct ที่เรียกใช้
	fmt.Println("discount price in struct pointer: ", f)

	fmt.Println("=======================================")

	zeroValuePointer()
}

func zeroValuePointer() {
	var c *course
	fmt.Printf("%#v\n", c)
	if c == nil {
		fmt.Println("c is nil")
	}
}

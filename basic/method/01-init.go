package main

import (
	"fmt"
)

type course struct {
	name, instructor string
	price            int
}

// method is a function that is associated with a type
// มันคือ function receiver ที่เกี่ยวข้องกับ type struct
// c is a receiver
func (c course) discount(d int) int {
	fmt.Printf("Discounting %s\n", c.price-d)
	return c.price - d
}

func (c course) info() {
	fmt.Printf("Name: %s\n", c.name)
	fmt.Printf("Instructor: %s\n", c.instructor)
	fmt.Printf("Price: %d\n", c.price)
}
func main() {
	c := course{
		name:       "Go",
		instructor: "Mihalis",
		price:      10000,
	}

	fmt.Printf("%#v\n", c)

	// can't do this because discount is a method not function
	// d := discount(c)
	d := c.discount(120)

	fmt.Printf("%#v\n", d)
	c.info()

	sum := 1
	for sum < 5 {
		sum += 1
	}
	fmt.Println(sum)
}

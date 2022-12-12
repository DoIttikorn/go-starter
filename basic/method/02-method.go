package main

import "fmt"

type curser interface {
	discount(d int) int
}

type course struct {
	price int
}

func (c course) discount(d int) int {
	return c.price - d
}

func main() {
	c := course{price: 10000}
	fmt.Println(c)
}

package main

import "fmt"

func main() {
	var msg string
	var age int
	var price float64
	var check bool
	var r rune

	fmt.Printf("msg: %s \n", msg)       // empty string
	fmt.Printf("age: %d \n", age)       // 0
	fmt.Printf("price: %.3f \n", price) // 0.000
	fmt.Printf("check: %t \n", check)   // false
	fmt.Printf("rune: %v \n", r)        // 0

}

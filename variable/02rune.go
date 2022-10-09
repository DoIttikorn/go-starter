package main

import "fmt"

func main() {
	var msg, age, price, check = "Hello", 10, 10.5, true
	var r rune = 'A'
	var a rune = 'Z'

	println(r)
	fmt.Printf("%c\n", a)

	fmt.Printf("msg: %s \n", msg)
	fmt.Printf("age: %d \n", age)
	fmt.Printf("price: %.3f \n", price)
	fmt.Printf("check: %t \n", check)
	fmt.Printf("rune: %v \n", r) // %v is value
}

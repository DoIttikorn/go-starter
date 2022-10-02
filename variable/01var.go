package main

import (
	"fmt"
)

// var price = 399.99 //  if have other file for use price so go run work

type ByteSize float64

// basic const run on compiler
// iota (incerate 1 always)
const (
	_           = iota
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

// basic var run on runtime

func main() {
	fmt.Println(TB)
	msg := "Nikato" + " Alias"
	var age, price, check = 20, 22.52, true
	// age, price, check := 20, 22.52, true

	fmt.Println("msg:", msg)
	fmt.Println("age:", age)
	fmt.Println("price:", price)
	fmt.Println("check:", check)
}

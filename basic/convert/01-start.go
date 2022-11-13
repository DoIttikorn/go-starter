package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 256
	fmt.Printf("type: %T, value: %v\n", i, i)

	var f float64 = float64(i)
	fmt.Printf("type: %T, value: %v\n", f, f)

	// var u uint = uint8(f) // cannot convert f (type float64) to type uint8
	// fmt.Printf("type: %T, value: %v\n", u, u)

	// convert a string to a number
	v := "72"

	s, err := strconv.Atoi(v)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("type: %T, value: %v\n", s, s)
}

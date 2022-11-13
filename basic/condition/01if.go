package main

import (
	"fmt"
	"math"
)

func main() {
	num := 33
	if num%2 == 0 {
		println("เลขคู่")
	} else if num%3 == 0 {
		println("เลขหาร 3 ลงตัว")
	} else {
		println("เลขคี่")
	}

	// short if //

	lim := 255.0
	// v declared in if statement
	if v := math.Pow(10, 2); v < lim {
		fmt.Println("x power n is: ", v)
	} else {
		fmt.Printf("x power n is %g over limit %g", v, lim)
	}
	// fmt.Println(v)

	// shop
	if rate := 8.4; rate < 5.0 {
		fmt.Println("Disappointing")
	} else if rate >= 5.0 && rate < 7.0 {
		fmt.Println("Normal")
	} else if rate >= 7.0 && rate < 10 {
		fmt.Println("Good")
	} else {
		fmt.Println("Excellent")
	}

}

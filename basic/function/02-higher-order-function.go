package main

import "fmt"

// func inc() int {
// 	return 1
// }

// func curr() int {
// 	return 2
// }

// higher order function
func adder() (func() int, func() int) {
	sum := 0
	return func() int {
			sum = sum + 1
			return sum
		}, func() int {
			return sum
		}
}

func main() {
	inc, curr := adder()
	v := inc()
	fmt.Println(v)
	v = inc()
	fmt.Println(v)
	v = inc()
	fmt.Println(v)
	v = curr()
	fmt.Println(v)
}

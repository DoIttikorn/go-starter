package main

import "fmt"

func main() {
	xs := []float64{4, 5, 7, 8, 3, 8, 0}
	ys := []float64{7, 2, 10, 9, 7}

	var votes []float64

	votes = append(xs, ys...)

	fmt.Printf("%#v\n", votes)
	fmt.Printf("%#v\n", votes[5:9])

}

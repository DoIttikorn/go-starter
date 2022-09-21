package main

import "fmt"

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
var (
	test = "TAI"
)

func main() {
	fmt.Println("test")
	fmt.Println(test)
}

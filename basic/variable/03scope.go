package main

var s = "Hello, playground"

func main() {
	println(s)
	x := true
	if x {
		y := 1
		if x != false {
			println(s)
			println(x)
			println(y)
		}
	}
	println(x)
}

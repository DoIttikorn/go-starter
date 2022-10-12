package main

import "fmt"

func show(sk []string) {
	l := len(sk)

	fmt.Printf("length: %d -- value: %#v", l, sk)
}

func main() {
	ss := make([]int, 3)
	// var ss []int
	// fmt.Printf("%#v\n", ss)

	// if ss != nil {
	// 	fmt.Println("not nil")
	// } else {
	// 	fmt.Println("nil")
	// }
	ss = append(ss, 33)
	fmt.Printf("%#v\n", ss)
}

package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	r := map[string]int{}
	for _, v := range strings.Fields(s) {
		r[v]++
	}
	return r
}

func main() {
	s := "If it look like a duck, swim like a duck, and quack like a duck, then it probably is a duck"
	w := WordCount(s)
	fmt.Printf("%#v\n", w)
}

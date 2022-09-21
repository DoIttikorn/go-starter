package main

import (
	"fmt"

	"doittikorn.dev/go-starter/imports"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(imports.Import2Banana("test"))
	fmt.Println(imports.Import2Coconut("test "))
}

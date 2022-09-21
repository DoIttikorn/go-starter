package main

import (
	"fmt"

	"doittikorn.dev/go-starter/imports"
	"doittikorn.dev/go-starter/variable"
)

func main() {
	fmt.Println("Hello, World!")
	variableMain()
}

// how to using variable
func variableMain() {
	variable.ExportShow()
}

// how to using import file
func importer() {
	fmt.Println(imports.Import2Banana("test"))
	fmt.Println(imports.Import2Coconut("test "))
}

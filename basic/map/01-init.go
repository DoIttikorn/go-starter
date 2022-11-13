package main

import "fmt"

// map is a data structure that associates a key with a value
func main() {
	// key is a string, value is an int
	var m map[string]int = map[string]int{}

	fmt.Printf("%#v\n", m)

	m["Answer"] = 42 // invalid operation: m['Answer'] (type map[string]int does not support indexing)

	fmt.Printf("%#v\n", m)

	// read a value form a map
	v := m["Nong"]
	fmt.Printf("%#v\n", v)

	m["Dodo"] = 10
	fmt.Printf("%#v\n", m)

	m["Dodo"] = 20
	fmt.Printf("%#v\n", m)

	// delete a key from a map
	delete(m, "Dodo")
	fmt.Printf("%#v\n", m)

	// try to read a key that does not exist
	n := m["Dodo"]
	fmt.Printf("%#v\n", n)

	// in go we can use the comma ok idiom to check if a key exists
	n, ok := m["Dodo"]
	if ok {
		fmt.Print("key exists")
		fmt.Print("The value is: ", n)
	} else {
		fmt.Print("key does not exist")
	}

	MakeMemoryMap()
}

// book memory
func MakeMemoryMap() {
	memoryMap := make(map[int]int)

	if memoryMap == nil {
		fmt.Println("memoryMap is nil")
	}

	fmt.Printf("memoryMap: %#v\n", memoryMap)
}

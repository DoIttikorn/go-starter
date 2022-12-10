package main

import "fmt"

// function literal คือ การประกาศฟังก์ชันแบบไม่มีชื่อ
func main() {
	r := func(a, b int) bool {
		return a < b
	}(2, 3)

	fmt.Println("result:", r)
}

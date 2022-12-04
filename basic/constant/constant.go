package main

import "fmt"

type day int

func main() {

	const (
		_          = iota
		Sunday day = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	fmt.Println("Sunday", Sunday)
	fmt.Println("Monday", Monday)
	fmt.Println("Tuesday", Tuesday)
	fmt.Println("Wednesday", Wednesday)
	fmt.Println("Thursday", Thursday)
	fmt.Println("Friday", Friday)
	fmt.Println("Saturday", Saturday)
	fmt.Printf("Saturday %T\n", Saturday) // ต่อให้เป็น type เดี่ยวกันแต่อาจจะไม่ได้ตัวเดียวกันเพราะอยู่คนละ package

}

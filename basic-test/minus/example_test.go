package minus

import "fmt"

// เราสามารถเขียน example ได้เพื่อให้เกิดการ test อัตโนมัติ

/*
func Example() {}
func ExampleFunc() {} // ตัวอย่างเช่น ExampleMinus
func ExampleType() {} // ตัวอย่างเช่น ExampleInt
func ExampleType_Method() {} // ตัวอย่างเช่น ExampleInt_Minus
*/

// godoc -http=:6060

// ExampleMinus คือ ตัวอย่างการใช้งานฟังก์ชัน Minus
func ExampleMinus() {
	got := Minus(10, 5)
	fmt.Println(got)
	// Output:
	// 5
}

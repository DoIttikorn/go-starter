package main

import (
	"fmt"
	"math"
)

// function as value

// signature ของฟังก์ชัน คือ ชื่อฟังก์ชัน และ พารามิเตอร์
func add(x, y float64) (float64, string) {
	return x + y, "good"
}

func swap(x, y string) (string, string) {
	return y, x
}

// Naked Return คือ การ return ค่าจากฟังก์ชันโดยไม่ต้องระบุชื่อตัวแปร
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var x string = "Hello World"
var y int = 42

// ฟังก์ชันที่มี signature คือ func(float64, float64) float64
// เราสามารถเปลี่ยนเป็น function ปกติได้แล้วยังทำงานเหมือนเดิมขอแค่มี signature ตรงกัน
var add2 func(float64, float64) float64 = func(x, y float64) float64 {
	return x + y
}

func compute(fn func(float64, float64) float64) float64 {
	fmt.Println("Compute")
	return fn(3, 4)
}

func hypot(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

func main() {
	a, b := add(5.0, 6.0)
	fmt.Println(a, b)

	x, y := swap("hello", "world")
	fmt.Println(x, y)

	fmt.Println(split(17))

	fmt.Printf("Type: %T Value: %v\n", add2, add2(5.0, 6.0))

	r := compute(add2)
	fmt.Println("r :", r)

	h := compute(hypot)
	fmt.Println("h :", h)
}

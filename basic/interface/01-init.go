package main

import "fmt"

// any คือ type ที่เป็น interface ที่ไม่มี method ใดๆ
// สามารถเก็บค่าทุก type ได้

func main() {
	// var v interface{}
	var v any // ตัวแปร v สามารถเก็บค่าได้ทุก type ที่เราสร้างขึ้นมา go เรียกว่า empty interface
	v = 36
	// fmt.Printf("%T %#v\n", v, v)

	show(v)
	showSwitch(v)

	v = "Hello"
	// fmt.Printf("%T %#v\n", v, v)

	show(v)
	showSwitch(v)

	showSwitch(35.3)
}

func show(val any) {
	// จะมีปัญหาถ้าเราไม่มีการ check type ก่อนใช้งาน
	i, check := val.(int) // แปลงค่า val ให้เป็น int
	if check {
		i = i + 2
		fmt.Println(i)
	}

	s, check := val.(string) // แปลงค่า val ให้เป็น string
	if check {
		fmt.Println(s)
	}
}

func showSwitch(val any) {
	switch v := val.(type) { // ตรวจสอบ type ของ val
	case int:
		i := v * 2
		fmt.Println(i)
	case string:
		s := v + " World"
		fmt.Println(s)
	default:
		fmt.Println("unknown type")
	}
}

package main

import "fmt"

type Decorator func(s string) error

func Use(next Decorator) Decorator {
	// ฟังก์ชันนี้จะถูกเรียกใช้งานก่อนฟังก์ชันที่ถูกส่งเข้ามา

	// ฟังก์ชันที่ถูกส่งเข้ามาจะถูกเรียกเมื่อ return ฟังก์ชันนี้ออกไป

	return func(c string) error {
		fmt.Println("do something before next")

		r := (c) + " should be here"
		return next(r)
	}
}

func home(s string) error {
	fmt.Println("home:", s)
	return nil
}

// higher order function คือ ฟังก์ชันที่รับฟังก์ชันเป็นพารามิเตอร์
// ใช้ในการเขียน middleware หรือ decorator
// เพื่อเพิ่มความสามารถให้กับฟังก์ชันที่ถูกส่งเข้ามา เช่น การเช็คสิทธิ์
func main() {
	wrapped := Use(home)

	w := wrapped("hello")

	fmt.Println("w:", w)
}

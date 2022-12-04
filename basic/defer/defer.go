package main

import "fmt"

func main() {
	// defer จะทำงานเมื่อ function ที่เรียกใช้มันเสร็จแล้ว
	// ถ้ามี defer หลายตัว จะทำงานจากล่างขึ้นบน
	// defer ทำการเป็น stack แบบ LIFO

	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		// จะทำการเก็บค่า i ไว้ใน memory แล้วเรียกใช้งาน
		// go จะ pass by value แต่เมื่อเป็น pointer จะ pass by reference
		defer func(n int) {
			fmt.Println(n)
		}(i)
		// ถ้าเป็นแบบนี้ i จะกลายเป็น 10 ทุกครั้ง เพราะเป็นการเรียกใช้งาน i ที่อยู่ในสุดท้ายของ function main
		// defer func() {
		// 	fmt.Println(i)
		// }()
	}
	fmt.Println("done")
}

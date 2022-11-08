package main

import "fmt"

func main() {
	var price int = 9999

	// ค่าที่เก็บอยู่เป็นที่อยู่ของ int ที่เก็บค่า 9999
	var addr *int = &price
	fmt.Printf("%T\n", addr)
	fmt.Println(price, addr)

	// addr = 9333 ไม่ได้เพราะ addr เป็น pointer ที่เก็บค่าเป็นที่อยู่ของ int ที่เก็บค่า 9999

	*addr = 9400 // wirte value to address บอก go ให้วิ่งไปที่ตำแหน่งที่คุณเก็บค่า 9999 แล้วเปลี่ยนค่าเป็น 9400

	fmt.Println(price, addr)

	v := *addr
	fmt.Println(v)

}

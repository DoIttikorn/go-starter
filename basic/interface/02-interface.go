package main

import "fmt"

// interface คือข้อตกลง ที่กำหนดให้ type ที่เข้ามาใช้งานต้องมี method ที่กำหนดไว้
// ถ้า type นั้นมี method ที่กำหนดไว้ ก็จะถูกเรียกว่า implement interface

// สิ่งที่ interface ต้องการต้องมี method ที่กำหนดไว้ มีมากกว่านั้นไม่เป็นไร
type promotion interface {
	discount() int
}

type course struct {
}

func (c course) discount() int {
	return 0
}
func (c course) info() {}

func sale(val promotion) {
	fmt.Printf("sale %#v\n", val.discount())
	// val.info() ไม่สามารถเรียกใช้ได้ เพราะไม่มี method info() ใน interface promotion
}

func main() {
	v := course{}

	v.info() // สามารถเรียกใช้ได้เพราะมี method info() ใน type course
	sale(v)
}

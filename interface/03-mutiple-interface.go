package main

import "fmt"

// interface คือข้อตกลง ที่กำหนดให้ type ที่เข้ามาใช้งานต้องมี method ที่กำหนดไว้
// ถ้า type นั้นมี method ที่กำหนดไว้ ก็จะถูกเรียกว่า implement interface

// สิ่งที่ interface ต้องการต้องมี method ที่กำหนดไว้ มีมากกว่านั้นไม่เป็นไร
// interface สามารถมี method มากกว่า 1 ตัวได้
type promotion interface {
	discount() int
}

type infoer interface {
	info()
}

// type ใดๆ สามารถ implement ได้หลาย interface ได้
type course struct {
}

func (c course) discount() int {
	return 0
}
func (c course) info() {
	fmt.Println("info", c)
}

func sale(val promotion) {
	fmt.Printf("sale %#v\n", val.discount())
	// val.info() ไม่สามารถเรียกใช้ได้ เพราะไม่มี method info() ใน interface promotion
}

func show(val infoer) {
	val.info()
}

// type presenter interface {
// 	discount() int
// 	info()
// }

// รวม interface ได้
type presenter interface {
	promotion
	infoer
}

func summary(val presenter) {
	fmt.Println(val.discount())
	val.info()
}

func main() {
	v := course{}

	v.info() // สามารถเรียกใช้ได้เพราะมี method info() ใน type course
	sale(v)

	show(v)

	summary(v)
}

// เราสามารถสร้าง struct ก่อนพอทำงานเสร็จแล้ว เราดูพฤติกรรมก่อนแล้วค่อยก็สร้าง interface ได้

// ในทางปฏิบัติ เราจะสร้าง interface ที่มี method น้อยๆเพื่อ reuse ได้ง่าย

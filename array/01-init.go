package main

import "fmt"

// ขนาดของ array จะต้องระบุไว้เสมอ
func show(skills [3]string) {
	fmt.Printf("len: %d, cap: %d, value: %#v\n", len(skills), cap(skills), skills)
}

func main() {
	// len 		// ความยาวของ array
	// cap 		// ความจุของ array
	// append 	// เพิ่มข้อมูลใน array
	// copy 		// คัดลอกข้อมูลจาก array อื่น
	// delete 	// ลบข้อมูลใน array
	// [...] 	// ระบุขนาดของ array โดยไม่ต้องระบุค่าใน array
	// [...]int{1, 2, 3} // ระบุขนาดของ array และกำหนดค่าใน array
	// [...]int{1: 1, 3: 3} // ระบุขนาดของ array และกำหนดค่าใน array โดยระบุ index ของ array

	// var skills [3]string = [3]string{"Go", "Python", "JavaScript"}
	skills := [3]string{"Go", "Python", "JavaScript"}

	show(skills)

	fmt.Println(skills[0])

}

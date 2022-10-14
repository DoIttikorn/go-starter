package main

import "fmt"

func show(tag string, sk []string) {
	l := len(sk)
	c := cap(sk)
	fmt.Printf("%s length: %d cap: %d -- value: %#v\n", tag, l, c, sk)
}
func main() {
	skills := []string{"Go", "Python", "JavaScript", "Java", "C++", "C#", "PHP", "Ruby", "Kotlin", "Swift"}

	s1 := skills[0:2]
	// cap = 3 เพราะว่ามันไม่ได้ทิ้งของตัวที่เหลือออกไปมันจะใช้ cap ของตัวที่เหลืออยู่แต่มันจะมองเป็น 2 ตัวแรก
	show("s1", s1)

	// ตัวที่อยู่ด้านหน้าตอน split go จะเอาออกไปไม่เอามาคิด cap แล้ว
	s2 := skills[1:3]
	show("s2", s2)

	s1[1] = "C++"

	show("s1", s1)
	show("s2", s2)
	// เราไม่สามารถ split เกิน capacity ได้
	// xx := skills[0:4]
	// show("xx", xx)

	// การเพิ่มข้อมูลใน slice
	// var newSkills  = skills[:]
	// var newSkills  = make([]string, 3)

}

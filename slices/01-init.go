package main

import "fmt"

func main() {
	// การสร้าง slice แบบมีค่าเริ่มต้น
	// skills := []string{"Go", "Python", "JavaScript"}
	// การสร้าง slice แบบไม่มีข้อมูล
	// การสร้าง slice แบบไม่กำหนดขนาด
	// ค่าเริ่มต้นของ slice จะเป็นค่า nil
	var skills []string
	// การเพิ่มข้อมูลใน slice
	skills = append(skills, "Go")
	skills = append(skills, "Python")
	skills = append(skills, "JavaScript")
	// การเพิ่มข้อมูลใน slice แบบไม่กำหนด index
	skills = append(skills, "C++", "C#", "Java")

	fmt.Printf("length: %d -- value: %#v\n", len(skills), skills)

	// การเพิ่มข้อมูลใน slice แบบกำหนด index
	skills = append(skills[:2], "PHP", "Ruby")

	fmt.Printf("length: %d -- value: %#v\n", len(skills), skills)

	// การลบข้อมูลใน slice
	skills = append(skills[:2], skills[3:]...)

	fmt.Printf("length: %d -- value: %#v\n", len(skills), skills)

	for i, value := range skills {
		println(i, value)
	}
	// การเข้าถึงข้อมูลใน slice
	println(skills[0])
	println(skills[1])
	println(skills[2])
	// println(skills[3])
	// println(skills[4])
	// println(skills[5])

	// println(skills[:])

}

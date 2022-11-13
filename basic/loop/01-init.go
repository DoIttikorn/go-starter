package main

import "fmt"

func main() {
	sum := 1

	// ตำแหน่งแรกคือการกำหนดค่าเริ่มต้น
	// ตำแหน่งที่สองคือเงื่อนไขการวนซ้ำ
	// ตำแหน่งที่สามคือการเพิ่มค่าในตัวแปรที่กำหนดไว้

	// while loop
	for sum < 10 {
		sum += sum
		println("sum =", sum)
	}
	fmt.Println("done sum =", sum)
	// การวนซ้ำแบบไม่มีเงื่อนไข
	for {
		println("loop")
		break
	}

	skills := []string{"Go", "Python", "JavaScript"}

	for i, v := range skills {
		fmt.Println(i, v)
	}
}

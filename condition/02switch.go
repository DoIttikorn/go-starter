package main

func main() {

	// fallthrough คือการทำให้ case ถัดไปทำงานต่อไป 1 ครั้ง
	// ถ้าไม่มี fallthrough จะหยุดทำงานทันที
	// short statement ใน case จะมี scope ใน case เท่านั้น
	switch today := "Saturday"; today {
	case "Monday":
		println("วันจันทร์")
	case "Tuesday":
		println("วันอังคาร")
	case "Wednesday":
		println("วันพุธ")
	case "Thursday":
		println("วันพฤหัสบดี")
	case "Friday":
		println("วันศุกร์")
	case "Saturday", "Sunday":
		println("วันหยุด")
		fallthrough
	default:
		println("ไม่มีวันนี้")
	}

	num := 1
	switch {
	case num < 0:
		println("nagative number")
	case num > 0:
		println("positive number")
	default:
		println("zero")
	}
	rating := 8.4
	switch {
	case rating < 5.0:
		println("bad")
	case rating >= 5.0 && rating < 7.0:
		println("normal")
	case rating >= 7.0 && rating < 9.0:
		println("good")
	case rating >= 9.0:
		println("excellent")
	default:
		println("unknown")
	}
}

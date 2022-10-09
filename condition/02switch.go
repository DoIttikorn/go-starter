package main

func main() {
	today := "Saturday"
	// fallthrough คือการทำให้ case ถัดไปทำงานต่อไป 1 ครั้ง
	switch today {
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
	case "Saturday":
		println("วันเสาร์")
		fallthrough
	case "Sunday":
		println("วันอาทิตย์")
	default:
		println("ไม่มีวันนี้")
	}
}

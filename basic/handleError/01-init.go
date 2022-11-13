package main

import "fmt"

func DivideError(a, b float64) float64 {
	return a / b
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Cannot divide by zero")
	}

	return a / b, nil
}

// ทำไม zero error ถึงเป็น nil เพราะ
/*
	type error interface {
		Error() string
	}
*/
func main() {
	r := DivideError(10, 0)
	fmt.Println(r)
	f, err := Divide(10, 0)
	if err != nil {
		fmt.Println("handle Error: ", err)
		return
	}
	fmt.Println(f)

}

package main

import "fmt"

func main() {
	// ขนาดของ array คือคุณสมบัติของ array ด้วย
	var genres [3]string = [3]string{"Action", "Adventure", "Fantasy"}

	genres[0] = "Action"
	genres[1] = "Adventure"
	genres[2] = "Fantasy"
	genres[1] = "sci-fi"

	fmt.Println(genres[0])
	fmt.Println(genres[1])
	fmt.Println(genres[2])

}

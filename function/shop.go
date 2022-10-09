package main

import "fmt"

var emote func(float64) string = func(ratings float64) string {
	if ratings < 5 {
		return "Disapponted 😔"
	} else if ratings >= 5 && ratings < 7 {
		return "Normal 😐"
	} else if ratings >= 7 && ratings < 10 {
		return "Good 😊"
	}
	return "👎"
}

func main() {

	fmt.Println(emote(4.5))
	fmt.Println(emote(5.5))
	fmt.Println(emote(7.5))
	fmt.Println(emote(10))
}

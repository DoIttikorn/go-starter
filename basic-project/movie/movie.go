package movie

import "fmt"

func init() {
	fmt.Println("init movie")
}

func Review(name string, rating float64) {
	fmt.Printf("I reviewed the movie %s with a rating of %f\n", name, rating)
}

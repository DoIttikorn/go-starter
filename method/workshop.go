package main

import "fmt"

type movie struct {
	title       string
	year        int
	rating      float32
	genres      []string
	isSuperHero bool
}

func (m movie) info() {
	fmt.Printf("Title: %s\n", m.title)
	fmt.Printf("Year: %d\n", m.year)
	fmt.Printf("Rating: %f\n", m.rating)
	fmt.Printf("Genres: %v	", m.genres)
}

func main() {
	m := movie{
		title:       "Batman",
		year:        2000,
		rating:      8.5,
		genres:      []string{"Action", "Adventure"},
		isSuperHero: true,
	}

	m.info()

}

package main

import "fmt"

type movie struct {
	title       string
	year        int
	rating      float64
	genres      []string
	isSuperHero bool
}

func main() {
	var m = movie{
		title:       "Avengers: Endgame",
		year:        2019,
		rating:      8.4,
		genres:      []string{"Action", "Adventure"},
		isSuperHero: true,
	}
	var m2 = movie{
		title:       "Batman",
		year:        2008,
		rating:      8.2,
		genres:      []string{"Action", "Adventure"},
		isSuperHero: true,
	}

	mvs := []movie{m, m2}

	// var mvs []movie = []movie{m, m2}

	// var mvs []movie
	// mvs = append(mvs, m, m2)

	for _, mv := range mvs {
		fmt.Printf("%+v\n", mv)
	}
}

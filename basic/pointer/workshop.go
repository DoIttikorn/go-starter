package main

import "fmt"

type movie struct {
	title       string
	year        int
	rating      float64
	votes       []float64
	genres      []string
	isSuperHero bool
}

func (m *movie) addVote(rating float64) {
	m.votes = append(m.votes, rating)
}

func main() {

	eg := &movie{
		title:       "The Matrix",
		year:        1999,
		rating:      8.7,
		votes:       []float64{8.7, 8.8, 8.9},
		genres:      []string{"Action", "Sci-Fi"},
		isSuperHero: false,
	}

	eg.addVote(8)
	fmt.Printf("eg.votes: %#+v\n", eg.votes)
}

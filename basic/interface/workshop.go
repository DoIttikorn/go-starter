package main

import "fmt"

type voter interface {
	addVote(float64)
}

func vote(v voter, rating float64) {
	v.addVote(rating)
}

type movie struct {
	title       string
	year        int
	rating      float32
	votes       []float64
	genres      []string
	isSuperHero bool
}

func (m *movie) addVote(rating float64) {
	m.votes = append(m.votes, rating)
}
func main() {

	eg := &movie{
		title:       "The Avengers",
		year:        2012,
		rating:      8.1,
		votes:       []float64{7.8, 8.1, 8.2, 8.3, 8.4},
		genres:      []string{"Action", "Adventure", "Sci-Fi"},
		isSuperHero: true,
	}

	vote(eg, 10)
	fmt.Println("Votes:", eg.votes)
}

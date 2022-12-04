package main

import (
	"encoding/json"
	"fmt"
)

type movie struct {
	Title       string   `json:"title"`
	Year        int      `json:"year"`
	Rating      float64  `json:"rating"`
	Geners      []string `json:"genres"`
	IsSuperHero bool     `json:"isSuperHero"`
}

func main() {
	body := []byte(`[
		{"title":"The Matrix","year":1999,"rating":8.7,"genres":["Action","Sci-Fi"],"isSuperHero":false},
		{"title":"The Matrix Reloaded","year":2003,"rating":7.2,"genres":["Action","Sci-Fi"],"isSuperHero":false}
	]`)

	// var m movie
	m := &[]movie{}
	err := json.Unmarshal(body, &m)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("m: %#+v\n", m)

}

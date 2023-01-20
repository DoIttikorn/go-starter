package main

import (
	"fmt"

	"github.com/Doittikorn/cinema/movie"
	"github.com/Doittikorn/cinema/ticket"
)

// init function is called sort by package name A-Z

// function is called if main package is executed
func init() {
	fmt.Println("init main")
}

func main() {
	movieName := movie.Find("tt0111161")
	fmt.Println(movieName)
	movie.Review(movieName, 9.2)
	ticket.BuyTicket(movieName)
}

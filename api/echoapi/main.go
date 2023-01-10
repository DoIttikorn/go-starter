package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Movie struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Year        int     `json:"year"`
	Rating      float64 `json:"rating"`
	IsSuperHero bool    `json:"isSuperHero"`
}

var movies = []Movie{
	{ID: 1, Title: "Batman", Year: 200, Rating: 8.5, IsSuperHero: true}}

func getAllMoviesHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, movies)
}
func main() {
	e := echo.New()

	e.GET("/movies", getAllMoviesHandler)
	port := 2565

	log.Println("Server started on port:", port)
	log.Fatal(e.Start(":" + strconv.Itoa(port)))

	// err := e.Start(":2565")
	// if err != nil {
	// 	log.Fatal(err)
	// }
}

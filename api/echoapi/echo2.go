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
	{ID: 1, Title: "Batman", Year: 2018, Rating: 8.5, IsSuperHero: true}}

func getAllMoviesHandler(c echo.Context) error {
	y := c.QueryParam("year")

	if y == "" {
		return c.JSON(http.StatusOK, movies)
	}

	year, err := strconv.Atoi(y)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid year",
		})
	}

	var filteredMovies []Movie
	for _, movie := range movies {
		if movie.Year == year {
			filteredMovies = append(filteredMovies, movie)
		}
	}

	if len(filteredMovies) == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Movie not found",
		})
	}
	return c.JSON(http.StatusOK, filteredMovies)
}

func getMovieByIdHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid id",
		})
	}
	for _, movie := range movies {
		if movie.ID == id {
			return c.JSON(http.StatusOK, movie)
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{
		"error": "Movie not found",
	})
}

func saveMovieHandler(c echo.Context) error {
	var movie Movie
	// err := c.Bind(&movie)
	// if err != nil {
	if err := c.Bind(&movie); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid payload",
		})
	}

	movie.ID = len(movies) + 1
	movies = append(movies, movie)

	return c.JSON(http.StatusCreated, movie)
}

func main() {
	e := echo.New()

	e.GET("/movies", getAllMoviesHandler)
	e.GET("/movies", saveMovieHandler)
	e.GET("/movies/:id", getMovieByIdHandler)
	port := 2565

	log.Println("Server started on port:", port)
	log.Fatal(e.Start(":" + strconv.Itoa(port)))

	// err := e.Start(":2565")
	// if err != nil {
	// 	log.Fatal(err)
	// }
}

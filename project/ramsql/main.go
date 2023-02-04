package main

import (
	"database/sql"
	"log"

	"github.com/labstack/echo/v4"
	_ "github.com/proullon/ramsql/driver"
)

type Movie struct {
	ImdbID      string
	Title       string
	Year        int
	Rating      float32
	IsSuperHero bool
}

var db *sql.DB

func con() {
	var err error
	db, err = sql.Open("ramsql", "goimdb")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

}
func main() {
	con()

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	// e.GET("/movies", getAllMoviesHandler)
	// e.GET("/movies/:id", getMovieHandler)
	// e.POST("/movies", addMovieHandler)

	e.Logger.Fatal(e.Start(":1323"))

}

package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/proullon/ramsql/driver"
)

type Movie struct {
	ID          int64   `json:"id"`
	ImdbID      string  `json:"imdbID"`
	Title       string  `json:"title"`
	Year        int     `json:"year"`
	Rating      float32 `json:"rating"`
	IsSuperHero bool    `json:"isSuperHero"`
}

var db *sql.DB

func updateMovieHandler(c echo.Context) error {
	imdbID := c.Param("imdbID")
	movie := new(Movie)
	// m := &Movie{}
	if err := c.Bind(movie); err != nil {
		return err
	}
	stmtCheck, err := db.Prepare(`
	SELECT id, imdbID, title, year, rating, isSuperHero
	FROM movies
	WHERE imdbID = ?;
	`)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer stmtCheck.Close()
	m := Movie{}
	err = stmtCheck.QueryRow(imdbID).Scan(&m.ID, &m.ImdbID, &m.Title, &m.Year, &m.Rating, &m.IsSuperHero)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if m.ID == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "No rows were returned!"})
	}

	stmt, err := db.Prepare(`
	UPDATE movies
	SET title=?, year=?, rating=?, isSuperHero=?
	WHERE imdbID=?;
	`)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer stmt.Close()
	b := fmt.Sprintf("%v", movie.IsSuperHero)
	_, err = stmt.Exec(movie.Title, movie.Year, movie.Rating, b, imdbID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	stmtQuery, err := db.Prepare(`
	SELECT id, imdbID, title, year, rating, isSuperHero
	FROM movies
	WHERE imdbID = ?;
	`)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer stmtQuery.Close()
	m = Movie{}
	err = stmtQuery.QueryRow(movie.ImdbID).Scan(&m.ID, &m.ImdbID, &m.Title, &m.Year, &m.Rating, &m.IsSuperHero)

	switch err {
	case nil:
		return c.JSON(http.StatusOK, m)
	case sql.ErrNoRows:
		return c.JSON(http.StatusNotFound, map[string]string{"message": "No rows were returned!"})
	default:
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
}

func getAllMoviesHandler(c echo.Context) error {
	year := c.QueryParam("year")
	movies := []Movie{}

	if year == "" {
		stmt, err := db.Prepare(`
	SELECT id, imdbID, title, year, rating, isSuperHero
	FROM movies;
	`)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		defer stmt.Close()
		rows, err := stmt.Query()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		defer rows.Close()
		for rows.Next() {
			var m Movie
			// m := Movie{}
			err := rows.Scan(&m.ID, &m.ImdbID, &m.Title, &m.Year, &m.Rating, &m.IsSuperHero)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, err.Error())
			}
			movies = append(movies, m)
		}

		if err := rows.Err(); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, movies)

	}
	years, err := strconv.Atoi(year)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	stmt, err := db.Prepare(`
	SELECT id, imdbID, title, year, rating, isSuperHero
	FROM movies
	WHERE year = ?;
	`)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer stmt.Close()
	rows, err := stmt.Query(years)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var m Movie
		// m := Movie{}
		err := rows.Scan(&m.ID, &m.ImdbID, &m.Title, &m.Year, &m.Rating, &m.IsSuperHero)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		movies = append(movies, m)
	}

	if err := rows.Err(); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, movies)

}

func getMovieByIdHandler(c echo.Context) error {
	imdbID := c.Param("imdbID")
	stmt, err := db.Prepare(`
	SELECT id, imdbID, title, year, rating, isSuperHero
	FROM movies
	WHERE imdbID = ?;
	`)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer stmt.Close()
	m := Movie{}
	err = stmt.QueryRow(imdbID).Scan(&m.ID, &m.ImdbID, &m.Title, &m.Year, &m.Rating, &m.IsSuperHero)
	switch err {
	case nil:
		return c.JSON(http.StatusOK, m)
	case sql.ErrNoRows:
		return c.JSON(http.StatusNotFound, map[string]string{"message": "No rows were returned!"})
	default:
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
}

func addMovieHandler(c echo.Context) error {
	m := new(Movie)
	// m := &Movie{}
	if err := c.Bind(m); err != nil {
		return err
	}
	stmt, err := db.Prepare(`
	INSERT INTO movies (imdbID, title, year, rating, isSuperHero) 
	VALUES (?, ?, ?, ?, ?);
	`)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer stmt.Close()
	b := fmt.Sprintf("%v", m.IsSuperHero)
	r, err := stmt.Exec(m.ImdbID, m.Title, m.Year, m.Rating, b)

	switch {
	case err == sql.ErrNoRows:
		return c.JSON(http.StatusNotFound, "No rows were returned!")
	// case err.Error() == "UNIQUE constraint violation":
	// 	return c.JSON(http.StatusConflict, "Movie already exists!")
	case err != nil:
		return c.JSON(http.StatusInternalServerError, err.Error())
	case err == nil:
		id, err := r.LastInsertId()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		m.ID = id
		return c.JSON(http.StatusCreated, m)
	default:
		return c.JSON(http.StatusInternalServerError, "Something went wrong!")
	}
}

func con() {
	var err error
	db, err = sql.Open("ramsql", "movies")
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

	createTable := `
	CREATE TABLE IF NOT EXISTS movies (
		id INT AUTO_INCREMENT PRIMARY KEY,
		imdbID TEXT NOT NULL UNIQUE,
		title TEXT NOT NULL,
		year INT NOT NULL,
		rating FLOAT NOT NULL,
		isSuperHero BOOLEAN
	);
	`
	_, err := db.Exec(createTable)
	if err != nil {
		log.Fatal("create table error", err)
	}

	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	e.GET("/movies", getAllMoviesHandler)
	e.GET("/movies/:imdbID", getMovieByIdHandler)
	e.POST("/movies", addMovieHandler)
	e.PUT("/movies/:imdbID", updateMovieHandler)

	e.Logger.Fatal(e.Start(":2565"))

}

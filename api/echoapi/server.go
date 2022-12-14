package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// เป็น struct ที่ implement ServeHTTP จะสามารถใช้งานกับ http.Handler ได้
type Logger struct {
	Handler http.Handler
}

var users = []User{
	{ID: 1, Name: "Ittikorn", Age: 30},
	{ID: 2, Name: "Ploy", Age: 30},
	{ID: 3, Name: "Wrap", Age: 30},
}

func healthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK")
}

type Err struct {
	Message string `json:"message"`
}

func createUsersHandler(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}
	users = append(users, *u)
	return c.JSON(http.StatusCreated, u)
}

func getUserHandler(c echo.Context) error {

	return c.JSON(http.StatusOK, users)
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/health", healthHandler)

	g := e.Group("/api")

	g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "ittikorn" && password == "1234" {
			return true, nil
		}
		return false, nil
	}))

	g.GET("/users", getUserHandler)
	g.POST("/users", createUsersHandler)

	log.Println("Listening on :2565")
	log.Fatal(e.Start(":2565"))
	log.Println("bye bye!!!")
}

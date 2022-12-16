package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/Doittikorn/ehoapisql/user"

	_ "github.com/lib/pq"
)

// เป็น struct ที่ implement ServeHTTP จะสามารถใช้งานกับ http.Handler ได้
type Logger struct {
	Handler http.Handler
}

func healthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK")
}

func main() {
	user.InitDB()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/health", healthHandler)

	e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "ittikorn" && password == "1234" {
			return true, nil
		}
		return false, nil
	}))

	e.GET("/users/:id", user.GetUserByIdHandler)
	e.GET("/users", user.GetUserHandler)
	e.POST("/users", user.CreateUsersHandler)

	log.Println("Listening on :2565")
	log.Fatal(e.Start(":2565"))
	log.Println("bye bye!!!")
}

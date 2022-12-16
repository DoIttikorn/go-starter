package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateUsersHandler(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}
	row := db.QueryRow("INSERT INTO users (name, age) VALUES ($1, $2) RETURNING id", u.Name, u.Age)

	// ขาดแค่ id ที่เรา insert ลงไป ให้เอามาเก็บไว้ในตัวแปร u.ID
	err := row.Scan(&u.ID)

	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, u)
}

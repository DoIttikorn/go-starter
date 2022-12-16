package user

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUserHandler(c echo.Context) error {
	stmt, err := db.Prepare("SELECT id,name,age FROM users")
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}
	rows, err := stmt.Query()

	users := []User{}

	for rows.Next() {
		u := User{}
		err := rows.Scan(&u.ID, &u.Name, &u.Age)

		if err != nil {
			return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
		}
		users = append(users, u)
	}

	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: "can't query all users:" + err.Error()})
	}
	return c.JSON(http.StatusOK, users)
}

func GetUserByIdHandler(c echo.Context) error {

	id := c.Param("id")
	// Prepare เป็น method ของ sql.DB ที่จะสร้าง query ขึ้นมา แล้ว return ค่าเป็น sql.Stmt
	stmt, err := db.Prepare("SELECT id,name,age FROM users WHERE id = $1") // ใช้ $1 แทนค่าที่จะส่งเข้าไปใน query
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't prepare query user statement " + err.Error()})
	}
	row := stmt.QueryRow(id) // QueryRow จะ return ค่าเป็น row ที่เป็น interface ที่มี method Scan ให้เราใช้

	u := User{}

	err = row.Scan(&u.ID, &u.Name, &u.Age)

	switch err {
	case sql.ErrNoRows:
		return c.JSON(http.StatusNotFound, Err{Message: "user not found"})
	case nil:
		return c.JSON(http.StatusOK, u)
	default:
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't scan user: " + err.Error()})
	}
}

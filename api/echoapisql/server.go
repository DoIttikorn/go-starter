package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/lib/pq"
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
	row := db.QueryRow("INSERT INTO users (name, age) VALUES ($1, $2) RETURNING id", u.Name, u.Age)

	// ขาดแค่ id ที่เรา insert ลงไป ให้เอามาเก็บไว้ในตัวแปร u.ID
	err := row.Scan(&u.ID)

	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, u)
}

func getUserHandler(c echo.Context) error {
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

func getUserByIdHandler(c echo.Context) error {

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

var db *sql.DB

func main() {
	e := echo.New()

	var err error
	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	createTb := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT,
		age INT
	);
	`
	_, err = db.Exec(createTb)

	if err != nil {
		log.Fatal("can't create table", err)
	}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/health", healthHandler)

	e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "ittikorn" && password == "1234" {
			return true, nil
		}
		return false, nil
	}))

	e.GET("/users/:id", getUserByIdHandler)
	e.GET("/users", getUserHandler)
	e.POST("/users", createUsersHandler)

	log.Println("Listening on :2565")
	log.Fatal(e.Start(":2565"))
	log.Println("bye bye!!!")
}

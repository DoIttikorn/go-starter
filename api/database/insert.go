package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	url := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", url)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	row := db.QueryRow("INSERT INTO users (name, email, age) VALUES ($1, $2, $3) RETURNING id", "Doit", "dodoittikorn@gmail.com", 20)

	var id int

	err = row.Scan(&id)

	if err != nil {
		panic(err)
	}
	fmt.Println("insert todo success id : ", id)
}

package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// ทำไมต้องมี _ หน้า github.com/lib/pq ครับ
// คือเราไม่ได้ใช้ package นี้ แต่เราต้อง import เพื่อให้เป็น dependency ของ package นี้
// func init ของ github.com/lib/pq จะ register driver ให้กับ database/sql ให้เรา

func init() {
	fmt.Println("main init")
}
func main() {
	url := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", url)

	if err != nil {
		panic(err)
	}

	defer db.Close()
	stmt, err := db.Prepare("select * from users")

	if err != nil {
		fmt.Println(err)
	}

	rows, err := stmt.Query()

	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		var id int
		var name string
		var email string
		var age int

		err = rows.Scan(&id, &name, &email, &age)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(id, name, email, age)
	}

	log.Println("query all success")
}

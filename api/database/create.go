package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/Doittikorn/gosql/dodo"
)

// ทำไมต้องมี _ หน้า github.com/lib/pq ครับ
// คือเราไม่ได้ใช้ package นี้ แต่เราต้อง import เพื่อให้เป็น dependency ของ package นี้
// func init ของ github.com/lib/pq จะ register driver ให้กับ database/sql ให้เรา

func init() {
	fmt.Println("main init")
}
func main() {
	dodo.Say()
	url := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", url)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	createTb := `
	CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name TEXT, email TEXT, age INT)
	`
	_, err = db.Exec(createTb)
	if err != nil {
		panic(err)
	}

	log.Println("create database success")
}

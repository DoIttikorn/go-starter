package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = []User{
	{ID: 1, Name: "Ittikorn", Age: 30},
	{ID: 2, Name: "Ploy", Age: 30},
	{ID: 3, Name: "Wrap", Age: 30},
}

func main() {
	http.HandleFunc("/users", func(w http.ResponseWriter, req *http.Request) {

		if req.Method == "GET" {
			b, err := json.Marshal(users)
			if err != nil {
				// check ว่ามี error หรือไม่
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)

				w.Write([]byte(err.Error()))
				return
			}
			w.Write(b)
			return
		}

		if req.Method == "POST" {
			// อ่านของจาก body ของ request
			body, err := ioutil.ReadAll(req.Body)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				fmt.Fprintf(w, "Error: %s", err)
				return
			}

			// แปลงข้อมูลจาก body ให้เป็น struct
			var u User
			err = json.Unmarshal(body, &u)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
				fmt.Fprintf(w, "Error: %s", err)
				return
			}

			users = append(users, u)

			fmt.Fprintf(w, "hello %s created users", "POST")
			// w.WriteHeader(http.StatusCreated)
			return
		}

		// set header error
		w.WriteHeader(http.StatusMethodNotAllowed)

	})

	log.Println("Listening on :2565")
	log.Fatal(http.ListenAndServe(":2565", nil))
	log.Println("bye bye!!!")
}

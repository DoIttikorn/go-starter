package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
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

func usersHandler(w http.ResponseWriter, req *http.Request) {
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

}

func healthHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func logMiddleware(Handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		Handler.ServeHTTP(w, r)
		log.Printf("Server http middleware : %s %s %s %s", r.Method, r.RequestURI, r.Proto, time.Since(start))
	}
}

// instance ใดๆ ที่ implement ServeHTTP จะสามารถใช้งานกับ http.Handler ได้
func (l Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) { // function นี้เป็นของ Logger ที่ implement ServeHTTP (receiver function)
	start := time.Now()
	l.Handler.ServeHTTP(w, r)
	log.Printf("Server http middleware : %s %s %s %s", r.Method, r.RequestURI, r.Proto, time.Since(start))
}

// Basic Auth
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, p, ok := r.BasicAuth()
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}
		if u != "dodo" || p != "123456" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Username or password is incorrect"))
			return
		}
		fmt.Println("AuthMiddleware")
		next(w, r)
	}

}

func main() {

	mux := http.NewServeMux()
	// เราจะใช้ logMiddleware กับ handler ที่เราสร้างขึ้นมา

	mux.HandleFunc("/users", AuthMiddleware(usersHandler))
	// http.HandleFunc("/users", logMiddleware(usersHandler))
	mux.HandleFunc("/health", healthHandler)
	// http.HandleFunc("/health", logMiddleware(healthHandler))

	logMux := Logger{Handler: mux}
	srv := http.Server{
		Addr:    ":2565",
		Handler: logMux,
	}
	log.Println("Listening on :2565")
	log.Fatal(srv.ListenAndServe())
	log.Println("bye bye!!!")
}

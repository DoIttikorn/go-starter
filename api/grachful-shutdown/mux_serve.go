package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("Dodo says: Hello, playground")

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Ittikorn"))
	})

	srv := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	go func() {
		log.Fatal(srv.ListenAndServe())
	}()

	fmt.Println("Server is running at port 8080")

	shutdown := make(chan os.Signal, 1)

	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	<-shutdown

	fmt.Println("Server is shutting down...")

	// ใช้ context.Background() ในการส่งค่า context ไปให้ Shutdown
	/*
		srv.Shutdown(context.Background())
		จะไม่รับ request ใหม่เข้ามาและจะทำการปิด connection ที่เปิดอยู่เมื่อส่งค่า context ไปให้ Shutdown
	*/
	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatal("Server failed to shutdown:", err)
	}

	fmt.Println("bye bye!")
}

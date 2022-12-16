package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	go func() {
		if err := e.Start(":1323"); err != nil && err != http.ErrServerClosed { // ต้อง check ด้วยว่าไม่ใช่ error ที่ main สั่งปิด err != http.ErrServerClosed
			e.Logger.Info("shutting down the server")
		}

	}()

	shutdown := make(chan os.Signal, 1)

	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	<-shutdown

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) // ต้องปิด server ให้เสร็จก่อน 10 วินาที
	defer cancel()

	if err := e.Shutdown(ctx); err != nil { // ถ้าไม่ปิด server ให้เสร็จก่อน 10 วินาที จะ error
		e.Logger.Fatal(err)
	}
}

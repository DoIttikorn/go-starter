package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	log.Println("server started")

	stop := make(chan os.Signal, 1)

	// os.Interrupt คือ event ที่เกิดขึ้นเมื่อมีการกด ctrl + c
	// syscall.SIGTERM คือ event ที่เกิดขึ้นเมื่อมีการปิดโปรแกรม
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	log.Println("wait for signal")
	<-stop
	log.Println("server stopped")
}

/*
	ใช้ channel ฟัง event ที่เกิดขึ้นใน os ได้
	ใช้ ps -ef | grep go หา pid ของ process ที่เราต้องการ
	ใช้ kill -9 pid ในการปิด process ที่เราต้องการ
	หรือ kill -SIGTINT pid เพื่อบอกว่าเป็นการ interrupt โปรแกรม
*/

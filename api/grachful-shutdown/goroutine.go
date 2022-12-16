package main

import (
	"fmt"
	"time"
)

func slow(s string) {
	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(s, ":", i)
	}
}

func main() {
	done := make(chan bool)
	go func() {
		slow("gopher")
		done <- true
	}()
	slow("python")
	<-done
	fmt.Println("all task done.")

	/*
		go slow("gopher") มันจะทำงานแบบ background และจะทำงานไปเรื่อยๆ จนกว่าจะมีการปิดโปรแกรม
		แต่ถ้าไม่มีตัว main ที่ทำงานอยู่ ตัว go จะทำงานแล้วจบไปเลย ไม่มีการทำงานอะไรเลย

		go slow("gopher")

	*/

	/*
		เหลือหากใส่ time.Sleep(5 * time.Second) ใน main จะทำให้โปรแกรมทำงานไป 5 วินาที แล้วจบ
		ถ้า go slow("gopher") ทำงานไป 10 วินาที โปรแกรมจะทำงานไป 5 วินาที แล้วจบ ไม่รอ go slow("gopher") ทำงานเสร็จ

		go slow("gopher")
		time.Sleep(5 * time.Second)

		ดังนั้นเราจะใช้ channel ในการรอ go slow("gopher") ทำงานเสร็จ
		channel คือ ตัวแปรที่สามารถส่งข้อมูลไปได้ และสามารถรับข้อมูลกลับมาได้ จะทำงานเหมือนท่อ
		channel จะทำงานเหมือนกับ time.Sleep แต่จะทำงานได้เร็วกว่า เพราะ time.Sleep จะทำงานแบบ blocking แต่ channel จะทำงานแบบ non-blocking

	*/
}

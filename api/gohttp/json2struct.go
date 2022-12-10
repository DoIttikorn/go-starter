package main

import (
	"encoding/json"
	"log"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	data := []byte(`{"id":1,"name":"Ittikorn","age":30}`)

	u := &User{}

	// Umarshal ต้องส่ง pointer ไปเพื่อเก็บค่าที่ได้จากการ Unmarshal
	// โดย struct ห้ามเป็น private ถ้าเป็น private จะไม่สามารถเข้าถึงได้
	err := json.Unmarshal(data, u)

	if err != nil {
		log.Fatal(err)
	}

	log.Fatalf("%#v", u)
}

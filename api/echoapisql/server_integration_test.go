//go:build integration

package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllUser(t *testing.T) {
	// ก่อนเราทำการ test เราต้องเพิ่มของใน database ก่อน
	seedUser(t)
	// ทดสอบการดึงข้อมูลทั้งหมด
	var us []User
	res := request(http.MethodGet, uri("users"), nil)
	err := res.Decode(&us) // ใช้ decode ข้อมูลที่ได้จาก response body มาเก็บไว้ในตัวแปร us

	// ตรวจสอบว่าไม่มี error
	assert.Nil(t, err)
	// ตรวจสอบว่า status code ต้องเท่ากับ 200
	assert.Equal(t, http.StatusOK, res.StatusCode)
	// ตรวจสอบว่าข้อมูลที่ได้มาจากการดึงข้อมูลทั้งหมดต้องมีค่ามากกว่า 0
	assert.Greater(t, len(us), 0)
}

func TestCrateUser(t *testing.T) {
	body := bytes.NewBufferString(`{
		"name":"dodo",
		"age": 23
	}`)
	var u User

	res := request(http.MethodPost, uri("users"), body)
	err := res.Decode(&u) // ใช้ decode ข้อมูลที่ได้จาก response body มาเก็บไว้ในตัวแปร u

	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, res.StatusCode)
	assert.NotEqual(t, 0, u.ID)
	assert.Equal(t, "dodo", u.Name)
	assert.Equal(t, 23, u.Age)

}

func TestGetUserById(t *testing.T) {
	c := seedUser(t)
	var latest User

	res := request(http.MethodGet, uri("users", strconv.Itoa(c.ID)), nil)
	err := res.Decode(&latest)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, c.ID, latest.ID)
	assert.NotEmpty(t, latest.Name)
	assert.NotEmpty(t, latest.Age)

}

func TestUpdateUser(t *testing.T) {
	t.Skip("TODO: implement me PATCH /users/:id")
}

func TestDeleteUser(t *testing.T) {
	t.Skip("TODO: implement me DELETE /users/:id")
}

func seedUser(t *testing.T) User {
	var c User
	body := bytes.NewBufferString(`{
		"name":"Ittikorn",
		"age": 23
	}`)

	err := request(http.MethodPost, uri("users"), body).Decode(&c)

	if err != nil {
		t.Fatal("can't create user:", err)
	}
	return c
}

type Response struct {
	*http.Response
	err error
}

// decode ของที่เราอยากจะได้จาก response
func (r *Response) Decode(v interface{}) error {
	if r.err != nil {
		return r.err
	}
	// ใช้ json.NewDecoder ในการ decode ข้อมูลจาก response body
	// เหมือนกับเราใช้ json.Unmarshal ในการ decode ข้อมูลจาก string
	return json.NewDecoder(r.Body).Decode(v)
}

// ใช้ในการยิง request ไปยัง server เพื่อทดสอบโดยใช้ http.NewRequest ในการสร้าง request
func request(method, url string, body io.Reader) *Response {
	req, _ := http.NewRequest(method, url, body)
	req.Header.Add("Authorization", os.Getenv("AUTH_TOKEN"))
	req.Header.Add("Content-Type", "application/json")
	client := http.Client{}
	resp, err := client.Do(req)
	return &Response{resp, err}
}

func uri(paths ...string) string {
	host := "http://localhost:2565"
	if paths == nil {
		return host
	}

	url := append([]string{host}, paths...)
	return strings.Join(url, "/")
}

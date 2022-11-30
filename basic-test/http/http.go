package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Response struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Info string `json:"info"`
}

var get = http.Get

// เป็น function ที่เรียกค่าจาก api แล้ว return ค่ากลับมา
// เวลา test เราจะ mock ค่าที่ return กลับมาโดยสร้าง server ขึ้นมาเอง
func MakeHTTPCall(url string) (*Response, error) {

	resp, err := get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	r := &Response{}

	if err := json.Unmarshal(body, r); err != nil {
		return nil, err
	}

	return r, nil

}

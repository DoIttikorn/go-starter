package http

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(`{"id": 1, "name": "Gopher", "info": "Golang"}`))
}

func TestMakeHttp(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(handler)) // สร้าง server ขึ้นมาเอง แล้วเรียกใช้ handler ที่เราสร้างไว้
	defer server.Close()

	// fmt.Println(server.URL)
	// time.Sleep(50 * time.Second)

	want := &Response{
		ID:   1,
		Name: "Gopher",
		Info: "Golang",
	}

	t.Run("Happy server response", func(t *testing.T) {

		resp, err := MakeHTTPCall(server.URL)

		if !reflect.DeepEqual(resp, want) { // DeepEqual เป็นการเปรียบเทียบค่าของ struct หรือ map หรือ slice หรือ array ที่มีค่าซ้ำกันหรือไม่
			t.Errorf("expected (%v) got (%v)", want, resp)
		}

		if !errors.Is(err, nil) {
			t.Errorf("expected (%v) got (%v)", nil, err)
		}
	})

}

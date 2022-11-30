package http

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

// func handler(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(http.StatusOK)

// 	w.Write([]byte(`{"id": 1, "name": "Gopher", "info": "Golang"}`))
// }

func TestMakeHttpCall(t *testing.T) {
	testTable := []struct {
		name             string
		server           *httptest.Server
		expectedResponse *Response
		expectedErr      error
	}{
		{
			name: "Happy server response",
			server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"id": 1, "name": "Gopher", "info": "Golang"}`))
			})),
			expectedResponse: &Response{
				ID:   1,
				Name: "Gopher",
				Info: "Golang",
			},
			expectedErr: nil,
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			defer tt.server.Close()

			resp, err := MakeHTTPCall(tt.server.URL)

			if !reflect.DeepEqual(resp, tt.expectedResponse) {
				t.Errorf("expected (%v) got (%v)", tt.expectedResponse, resp)
			}

			if !errors.Is(err, tt.expectedErr) {
				t.Errorf("expected (%v) got (%v)", tt.expectedErr, err)
			}
		})
	}
}

package main

import (
	"net/http"
	"testing"
)

const (
	GET  = "GET"
	POST = "POST"
	DEL  = "DEL"
)

func TestAPI(t *testing.T) {
	// type payload struct {
	// 	Word string `json:"word"`
	// }
	// payload := payload{
	// 	Word: "hello",
	// }
	tests := []struct {
		method         string
		endpoint       string
		expectedStatus int
	}{
		{GET, "/", 200},
		{GET, "/new", 200},
		{GET, "/api/v1/grids/201", 200},
		{POST, "/api/v1/words", 200},
		{DEL, "/api/v1/words", 200},
	}
	for i := range tests {
		switch tests[i].method {
		case GET:
			actual, _ := http.Get(tests[i].endpoint)
			if actual.StatusCode != tests[i].expectedStatus {
				t.Errorf("%s %s, expected=%v, received=%v", tests[i].method, tests[i].endpoint, tests[i].expectedStatus, actual.StatusCode)
				return
			}
			// case POST:
			// 	actual, _ := http.Post(tests[i].endpoint, payload)
			// 	if actual.StatusCode != tests[i].expectedStatus {
			// 		t.Errorf("%s %s, expected=v, received=%v", tests[i].method, tests[i].endpoint, tests[i].expectedStatus, actual.StatusCode)
			// 		return
			// 	}
		}
	}
}

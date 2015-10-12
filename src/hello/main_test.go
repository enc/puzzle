package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetRequest(t *testing.T) {
	s := NewServer()

	request, _ := http.NewRequest("GET", "/", nil)
	request.Header.Set("Accept-Language", "de-DE,de;q=0.8,en-US;q=0.6,en;q=0.4")
	response := httptest.NewRecorder()
	s.Respond(response, request)
	if body := response.Body.String(); !strings.Contains(body, "Your language is: de-DE") ||
		!strings.Contains(body, "You sent a: GET") {
		t.Fatalf("Expected language and method in body, got: %q\n", body)
	}

	request, _ = http.NewRequest("GET", "/", nil)
	request.Header.Set("Accept-Language", "en-US,en;q=0.8,de-DE;q=0.6,en;q=0.4")
	response = httptest.NewRecorder()
	s.Respond(response, request)
	if body := response.Body.String(); !strings.Contains(body, "Your language is: en-US") ||
		!strings.Contains(body, "You sent a: GET") {
		t.Fatalf("Expected language and method in body, got: %q\n", body)
	}

}

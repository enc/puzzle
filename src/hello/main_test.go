package main

import (
	"bytes"
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
	s.ServeHTTP(response, request)
	if body := response.Body.String(); !strings.Contains(body, "Your language is: de") ||
		!strings.Contains(body, "You sent a: GET") {
		t.Fatalf("Expected language and method in body, got: %q\n", body)
	}

	request, _ = http.NewRequest("GET", "/", nil)
	request.Header.Set("Accept-Language", "en-US,en;q=0.8,de-DE;q=0.6,en;q=0.4")
	response = httptest.NewRecorder()
	s.ServeHTTP(response, request)
	if body := response.Body.String(); !strings.Contains(body, "Your language is: en") ||
		!strings.Contains(body, "You sent a: GET") {
		t.Fatalf("Expected language and method in body, got: %q\n", body)
	}

}
func TestLanguageNotSet(t *testing.T) {
	s := NewServer()

	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()
	s.ServeHTTP(response, request)
	if body := response.Body.String(); !strings.Contains(body, "Your language is: Not Available") ||
		!strings.Contains(body, "You sent a: GET") {
		t.Fatalf("Expected language and method in body, got: %q\n", body)
	}
}

func TestPostRequest(t *testing.T) {
	formString := []byte("postVar=HalloWorld")

	s := NewServer()
	request, _ := http.NewRequest("POST", "/", bytes.NewBuffer(formString))
	request.Header.Set("Accept-Language", "en-US,en;q=0.8,de-DE;q=0.6,en;q=0.4")
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response := httptest.NewRecorder()
	s.ServeHTTP(response, request)
	if body := response.Body.String(); !strings.Contains(body, "Your POST variable value: HalloWorld") ||
		!strings.Contains(body, "You sent a: POST") {
		t.Fatalf("Missing post message. Got: %q\n", body)
	}
}
func TestEmptyPostRequest(t *testing.T) {
	s := NewServer()
	request, _ := http.NewRequest("POST", "/", nil)
	request.Header.Set("Accept-Language", "en-US,en;q=0.8,de-DE;q=0.6,en;q=0.4")
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response := httptest.NewRecorder()
	s.ServeHTTP(response, request)
	if body := response.Body.String(); response.Code != 400 ||
		!strings.Contains(body, "forgotten the postVar") {
		t.Fatalf("Expected 400 error but got : %d.\n", response.Code)
	}
}

func TestWrongUri(t *testing.T) {
	s := NewServer()
	request, _ := http.NewRequest("GET", "/faber", nil)
	response := httptest.NewRecorder()
	s.ServeHTTP(response, request)
	if response.Code != http.StatusNotFound {
		t.Fatalf("Expected StatusCode 404 but got: %d\n", response.Code)
	}
}

func TestPutRequest(t *testing.T) {
	s := NewServer()
	request, _ := http.NewRequest("PUT", "/", nil)
	response := httptest.NewRecorder()
	s.ServeHTTP(response, request)
	if response.Code != http.StatusMethodNotAllowed {
		t.Fatalf("Expected StatusCode 405 but got: %d\n", response.Code)
	}
}

func TestLanguageParsing(t *testing.T) {
	var tests = []struct {
		ls  string
		out string
	}{
		{"", "Not Available"},
		{"en-US,en;q=0.8,de-DE;q=0.6,en;q=0.4", "en"},
		{"de-DE,de;q=0.8,en-US;q=0.6,en;q=0.4", "de"},
		{"fr, de-DE,de;q=0.8,en-US;q=0.6,en;q=0.4", "fr"},
	}

	s := NewServer()
	for _, test := range tests {
		if out := s.ParseLanguage(test.ls); test.out != out {
			t.Fatalf("Parsed %s and expected %s but got: %s\n", test.ls, test.out, out)
		}
	}
}

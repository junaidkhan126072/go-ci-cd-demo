package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	// Step 1: Create a fake HTTP request
	req := httptest.NewRequest("GET", "/", nil)

	// Step 2: Create a ResponseRecorder to record the response
	rec := httptest.NewRecorder()

	// Step 3: Call the handler
	HelloHandler(rec, req)

	// Step 4: Check the response body
	expected := "🚀 Hello from CI/CD! kaka Boom\n"
	actual := rec.Body.String()

	if strings.TrimSpace(actual) != strings.TrimSpace(expected) {
		t.Errorf("Expected %q but got %q", expected, actual)
	}

	// Optional: check the HTTP status code
	if status := rec.Code; status != http.StatusOK {
		t.Errorf("Expected status 200, got %d", status)
	}
}

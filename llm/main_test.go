package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSetupRouter(t *testing.T) {
	mux := setupMux()

	req := httptest.NewRequest(http.MethodGet, "/llms", nil)

	rr := httptest.NewRecorder()

	mux.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	bodyBytes, err := io.ReadAll(rr.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}
	actualBody := string(bodyBytes)

	expectedBody := "LLM service"
	if actualBody != expectedBody {
		t.Errorf("handler returned unexpected body: got %v want %v",
			actualBody, expectedBody)
	}
}

func TestNotFound(t *testing.T) {
	router := setupMux()
	req := httptest.NewRequest(http.MethodGet, "/nonexistent", nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code for non-existent path: got %v want %v",
			status, http.StatusNotFound)
	}
}

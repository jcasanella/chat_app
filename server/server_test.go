package server

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jcasanella/chat_app/repository"
)

func TestLoginInvalid(t *testing.T) {
	expected := `{"error":"Invalid user"}`

	st := repository.NewMemStorage()
	s := NewServer(st)
	r := s.newRouter()
	req, _ := http.NewRequest("GET", "/v1/login", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)
	actual := string(responseData)

	if actual != expected {
		t.Errorf("NewRouter() /v1/login --> Actual response: %s Expected response: %s", actual, expected)
	}
	if w.Code != http.StatusBadRequest {
		t.Errorf("NewRouter() /v1/login --> Actual status code: %v Expected status code: %v", w.Code, http.StatusBadRequest)
	}
}

func TestLoginValid(t *testing.T) {
	// TODO
}

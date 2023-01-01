package server

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestLoginInvalid(t *testing.T) {
	expected := `{"error":"Invalid user"}`
	r := NewRouter()
	req, _ := http.NewRequest("GET", "/v1/login/", nil)
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

func getUserPayload() string {
	params := url.Values{}
	params.Add("name", "u1")
	params.Add("password", "p1")

	return params.Encode()
}

func TestLoginValid(t *testing.T) {
	// TODO
}

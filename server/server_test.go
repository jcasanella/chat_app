package server

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/jcasanella/chat_app/model"
	"github.com/jcasanella/chat_app/repository"
)

func initNewRouter() *gin.Engine {
	st := repository.NewMemStorage()
	s := NewServer(st)
	return s.newRouter()
}

func doRequest(payload io.Reader) (w *httptest.ResponseRecorder) {
	r := initNewRouter()

	req, _ := http.NewRequest("GET", "/v1/login", payload)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	return
}

func TestLoginInvalid(t *testing.T) {
	w := doRequest(nil)

	responseData, _ := io.ReadAll(w.Body)
	actual := string(responseData)

	expected := `{"error":"invalid request"}`
	if actual != expected {
		t.Errorf("NewRouter() /v1/login --> Actual response: %s Expected response: %s", actual, expected)
	}
	if w.Code != http.StatusBadRequest {
		t.Errorf("NewRouter() /v1/login --> Actual status code: %v Expected status code: %v", w.Code, http.StatusBadRequest)
	}
}

func createValidUser() bytes.Buffer {
	var buf bytes.Buffer

	u := model.User{Name: "admin", Password: "password"}
	err := json.NewEncoder(&buf).Encode(u)
	if err != nil {
		log.Fatal(err)
	}

	return buf
}

func TestLoginValid(t *testing.T) {
	buf := createValidUser()
	w := doRequest(&buf)
	responseData, _ := io.ReadAll(w.Body)
	actual := string(responseData)

	expected := `{"token":`
	if !strings.Contains(actual, expected) {
		t.Errorf("NewRouter() /v1/login --> Actual response: %s Expected response: %s", actual, expected)
	}
	if w.Code != http.StatusOK {
		t.Errorf("NewRouter() /v1/login --> Actual status code: %v Expected status code: %v", w.Code, http.StatusOK)
	}
}

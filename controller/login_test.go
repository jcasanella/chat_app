package controller

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func getTestContext(w *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)

	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL: &url.URL{
			Host: "localhost:8080",
			Path: "/v1/login",
		},
	}

	return ctx
}

func mockGetJSON(c *gin.Context, user UserTest) {
	c.Request.Method = "GET"
	c.Request.Header.Set("Content-Type", "application/json")

	b, _ := json.Marshal(user)
	c.Request.Body = io.NopCloser(bytes.NewBuffer(b))
}

func TestValidLogin(t *testing.T) {
	expected := `{"token":`

	lc := new(LoginController)
	w := httptest.NewRecorder()

	c := getTestContext(w)

	user := &UserTest{Name: "Frank", Password: "p1"}
	mockGetJSON(c, *user)

	lc.Login(c)

	resp, _ := io.ReadAll(w.Body)
	actual := string(resp)
	if !strings.Contains(actual, expected) {
		t.Errorf("Login() --> Actual response: %s does not contain Expected response: %s", actual, expected)
	}
	if w.Code != http.StatusOK {
		t.Errorf("Login() --> Actual status code: %v Expected status code: %v", w.Code, http.StatusOK)
	}
}

func createURLValues() []UserTest {
	// Wrong args
	values1 := &UserTest{}

	// Only name
	values2 := &UserTest{Name: "peter"}

	// Only password
	values3 := &UserTest{Password: "pass"}

	return []UserTest{*values1, *values2, *values3}
}

func TestInvalidLogin(t *testing.T) {
	expected := `{"error":"Invalid user"}`

	for _, v := range createURLValues() {
		lc := new(LoginController)
		w := httptest.NewRecorder()

		c := getTestContext(w)

		mockGetJSON(c, v)

		lc.Login(c)

		resp, _ := io.ReadAll(w.Body)
		if string(resp) != expected {
			t.Errorf("Login() --> Actual response: %s Expected response: %s", string(resp), expected)
		}
		if w.Code != http.StatusBadRequest {
			t.Errorf("Login() --> Actual status code: %v Expected status code: %v", w.Code, http.StatusBadRequest)
		}
	}
}

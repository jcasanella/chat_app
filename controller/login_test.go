package controller

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
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

func mockGetJson(c *gin.Context, params gin.Params, u url.Values) {
	c.Request.Method = "GET"
	c.Request.Header.Set("Content-Type", "text/plain")

	c.Params = params

	c.Request.URL.RawQuery = u.Encode()
}

func TestValidLogin(t *testing.T) {
	expected := "Valid user"

	lc := new(LoginController)
	w := httptest.NewRecorder()

	c := getTestContext(w)

	values := url.Values{}
	values.Add("name", "user1")
	values.Add("password", "pass1")

	mockGetJson(c, nil, values)

	lc.Login(c)

	resp, _ := io.ReadAll(w.Body)
	if string(resp) != expected {
		t.Errorf("Login() --> Actual response: %s Expected response: %s", string(resp), expected)
	}
	if w.Code != http.StatusOK {
		t.Errorf("Login() --> Actual status code: %v Expected status code: %v", w.Code, http.StatusOK)
	}
}

func createUrlValues() []url.Values {
	// Wrong args
	values := url.Values{}
	values.Add("wrongName", "wrong")
	values.Add("wrongPassword", "wrong")

	// Only name
	values2 := url.Values{}
	values2.Add("name", "peter")

	// Only password
	values3 := url.Values{}
	values3.Add("password", "pass")

	return []url.Values{values, values2, values3}
}

func TestInvalidLogin(t *testing.T) {
	expected := "Invalid user"

	for _, v := range createUrlValues() {
		lc := new(LoginController)
		w := httptest.NewRecorder()

		c := getTestContext(w)

		mockGetJson(c, nil, v)

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

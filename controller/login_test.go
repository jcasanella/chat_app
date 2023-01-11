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
	"github.com/jcasanella/chat_app/model"
	"github.com/jcasanella/chat_app/repository"
	"github.com/jcasanella/chat_app/service"
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

func mockGetJSON(c *gin.Context, user model.User) {
	c.Request.Method = "GET"
	c.Request.Header.Set("Content-Type", "application/json")

	b, _ := json.Marshal(user)
	c.Request.Body = io.NopCloser(bytes.NewBuffer(b))
}

func initUserService() *service.UserService {
	st := repository.NewMemStorage()
	db := repository.NewServiceDb(st)
	return service.NewUserService(db)
}

func TestValidLogin(t *testing.T) {
	expected := `{"token":`

	us := initUserService()
	lc := NewLoginController(us)
	w := httptest.NewRecorder()
	c := getTestContext(w)

	user := &model.User{Name: "admin", Password: "password"}
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

func createURLValues() []model.User {
	// Wrong args
	values1 := &model.User{}

	// Only name
	values2 := &model.User{Name: "peter"}

	// Only password
	values3 := &model.User{Password: "pass"}

	return []model.User{*values1, *values2, *values3}
}

func TestInvalidLogin(t *testing.T) {
	expected := `{"error":"Key:`

	for _, v := range createURLValues() {
		us := initUserService()
		lc := NewLoginController(us)
		w := httptest.NewRecorder()
		c := getTestContext(w)

		mockGetJSON(c, v)

		lc.Login(c)

		resp, _ := io.ReadAll(w.Body)
		s := string(resp)
		if !strings.Contains(s, expected) {
			t.Errorf("Login() --> Actual response: %s Expected response: %s", string(resp), expected)
		}
		if w.Code != http.StatusBadRequest {
			t.Errorf("Login() --> Actual status code: %v Expected status code: %v", w.Code, http.StatusBadRequest)
		}
	}
}

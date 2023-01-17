package service

import (
	"testing"

	"github.com/jcasanella/chat_app/model"
	"github.com/jcasanella/chat_app/repository"
)

func initUserService() *UserService {
	st := repository.NewMemStorage()
	db := repository.NewServiceDb(st)
	return NewUserService(db)
}

func TestGetInvalidUser(t *testing.T) {
	invalidUsers := []model.User{
		{},
		{Name: "p"},
		{Password: "p"},
		{Name: "name", Password: "password"},
	}

	us := initUserService()
	for _, v := range invalidUsers {
		u, err := us.GetUser(v)
		if u != nil {
			t.Errorf("GetUser()  --> Expected nil user and actual: %v", u)
		}
		if err == nil {
			t.Errorf("GetUser()  --> Expected error and actual is nil")
		}
	}
}

func TestGetValidUser(t *testing.T) {
	validUser := model.User{Name: "admin", Password: "password"}
	us := initUserService()

	u, err := us.GetUser(validUser)
	if *u != validUser {
		t.Errorf("GetUser()  --> Expected user: %v actual: %v", *u, validUser)
	}

	if err != nil {
		t.Errorf("GetUser()  --> Expected error nil actual: %v", err)
	}
}

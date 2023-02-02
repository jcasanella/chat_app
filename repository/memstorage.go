package repository

import (
	"fmt"

	"github.com/jcasanella/chat_app/model"
)

type MemStorage struct {
	Users map[string]string
}

// NewMemStorage initialize a MemStorage
func NewMemStorage() MemStorage {
	ms := MemStorage{
		Users: make(map[string]string),
	}

	ms.Users["admin"] = "password"
	return ms
}

// Get receives a key and a password and look in the repository for the user, if exists return the user otherwise error
func (m MemStorage) Get(key string, password string) (*model.User, error) {
	p := m.Users[key]
	if p == "" {
		return nil, fmt.Errorf("error user does not exist")
	}

	return &model.User{Name: key, Password: password}, nil
}

// Add receives a key and a password, if the user is added to the repository returns the user, otherwise error
func (m MemStorage) Add(key string, password string) (*model.User, error) {
	p := m.Users[key]
	if p != "" {
		return nil, fmt.Errorf("error user exists")
	}

	m.Users[key] = password

	return &model.User{Name: key, Password: password}, nil
}

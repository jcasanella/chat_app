package repository

import (
	"fmt"

	"github.com/jcasanella/chat_app/model"
)

type MemStorage struct {
	Users map[string]string
}

func NewMemStorage() MemStorage {
	ms := MemStorage{
		Users: make(map[string]string),
	}

	ms.Users["admin"] = "password"
	return ms
}

func (m MemStorage) Get(key string, password string) (*model.User, error) {
	p := m.Users[key]
	if p == "" {
		return nil, fmt.Errorf("error user does not exist")
	}

	return &model.User{Name: key, Password: password}, nil
}

func (m *MemStorage) Add(key string, password string) (*model.User, error) {
	p := m.Users[key]
	if p != "" {
		return nil, fmt.Errorf("error user exists")
	}

	m.Users[key] = password

	return &model.User{Name: key, Password: password}, nil
}

package repository

import (
	"fmt"

	"github.com/jcasanella/chat_app/model"
)

type MemStorage struct {
	Users map[string]string
}

func NewMemStorage() *MemStorage {
	ms := &MemStorage{
		Users: make(map[string]string),
	}

	ms.Users["admin"] = "password"
	return ms
}

func (m MemStorage) Get(user model.User) (model.User, error) {
	p := m.Users[user.Name]
	if p == "" || p != user.Password {
		return model.User{}, fmt.Errorf("error User does not exists")
	}

	return user, nil
}

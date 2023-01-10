package repository

import (
	"fmt"
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

func (m MemStorage) Get(key string) (string, error) {
	p := m.Users[key]
	if p == "" {
		return "", fmt.Errorf("error User does not exists")
	}

	return p, nil
}

package repository

import "github.com/jcasanella/chat_app/model"

type Storage interface {
	Get(user model.User) (model.User, error)
}

type ServiceDB struct {
	MyDB Storage
}

func NewServiceDb(db Storage) *ServiceDB {
	return &ServiceDB{
		MyDB: db,
	}
}
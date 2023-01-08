package repository

import "github.com/jcasanella/chat_app/model"

type Storage interface {
	Get(user model.User) (model.User, error)
}

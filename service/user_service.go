package service

import (
	"github.com/jcasanella/chat_app/model"
	"github.com/jcasanella/chat_app/repository"
)

type UserService struct {
	s *repository.ServiceDB
}

func NewUserService(rsDb *repository.ServiceDB) *UserService {
	return &UserService{
		s: rsDb,
	}
}

func (u UserService) Get(user model.User) {
	p, err := u.s.MyDB.Get(user.Name)
	if err != nil {

	} else {
		if p == user.Password {

		} else {

		}
	}
}

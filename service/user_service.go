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

func (us UserService) GetUser(user model.User) (*model.User, error) {
	if p, err := us.s.MyDB.Get(user.Name, user.Password); err != nil {
		return nil, err
	} else {
		return p, nil
	}
}

func (us *UserService) AddUser(user model.User) (*model.User, error) {
	if p, err := us.s.MyDB.Add(user.Name, user.Password); err != nil {
		return nil, err
	} else {
		return p, nil
	}
}

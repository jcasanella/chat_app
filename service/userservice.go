package service

import (
	"github.com/jcasanella/chat_app/model"
	"github.com/jcasanella/chat_app/repository"
)

type UserService struct {
	s *repository.ServiceDB
}

// NewUserService initialize an UserService
func NewUserService(rsDb *repository.ServiceDB) *UserService {
	return &UserService{
		s: rsDb,
	}
}

// GetUser receives an user and look for the user in the database, if exists return the user, otherwise error
func (us UserService) GetUser(user model.User) (*model.User, error) {
	if p, err := us.s.MyDB.Get(user.Name, user.Password); err != nil {
		return nil, err
	} else {
		return p, nil
	}
}

// AddUser receives an user and adds them into the database, if added returns the user, otherwise error
func (us *UserService) AddUser(user model.User) (*model.User, error) {
	if p, err := us.s.MyDB.Add(user.Name, user.Password); err != nil {
		return nil, err
	} else {
		return p, nil
	}
}

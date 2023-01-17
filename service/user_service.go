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

func (u UserService) GetUser(user model.User) (*model.User, error) {
	if p, err := u.s.MyDB.Get(user.Name, user.Password); err != nil {
		return nil, err
	} else {
		return p, nil
	}
}

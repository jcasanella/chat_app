package service

import (
	"fmt"

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

func (u UserService) Get(user model.User) (*model.User, error) {
	if p, err := u.s.MyDB.Get(user.Name); err != nil {
		return nil, err
	} else if p == user.Password {
		return &user, nil
	} else {
		return nil, fmt.Errorf("error invalid password")
	}
}

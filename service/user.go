package service

import (
	"errors"
	"final-project/helper"
	"final-project/model"
)

type UserService interface {
	Register(user model.User) (model.User, error)
	Login(model.LoginCredentials) (string, error)
	Validate(model.User) error
}

func (s *Service) Register(user model.User) (model.User, error) {
	user.Password = helper.HashPass(user.Password)
	return s.repo.Register(user)
}

func (s *Service) Login(user model.LoginCredentials) (string, error) {
	res, isAvail, _ := s.repo.FindByEmail(user.Email)

	if !isAvail {
		return "", errors.New("invalid email / password")
	}

	comparedPassword := helper.ComparePass(res.Password, user.Password)

	if !comparedPassword {
		return "", errors.New("invalid password")
	}

	token := helper.GenerateToken(uint(res.ID), res.Email)

	return token, nil

}

func (s *Service) Validate(user model.User) error {
	_, usernameAvail, _ := s.repo.FindByUsername(user.UserName)
	_, emailAvail, _ := s.repo.FindByEmail(user.Email)

	if usernameAvail && emailAvail {
		return errors.New("username and email already exist")
	} else if usernameAvail {
		return errors.New("username already exist")
	} else if emailAvail {
		return errors.New("email already exist")
	}

	return nil
}

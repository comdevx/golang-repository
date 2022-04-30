package service

import (
	"errors"
	logs "project/helper"
	"project/repository"

	"golang.org/x/crypto/bcrypt"
)

type authenService struct {
	userRepo repository.UserRepository
}

func NewAuthenService(userRepo repository.UserRepository) authenService {
	return authenService{userRepo: userRepo}
}

func (s authenService) Login(body AuthenBody) (*AuthenResponse, error) {

	user, err := s.userRepo.GetByUser(body.Username)
	if err != nil {
		logs.Error(err)
		return nil, ErrServerError()
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		return nil, errors.New("Password is not matched")
	}

	if user.Username == body.Username && user.Password == "" {
		return nil, ErrValidationError("Invalid username or password")
	}

	token, err := logs.GenerateToken(user.Username, user.ID)
	if err != nil {
		return nil, err
	}

	response := &AuthenResponse{}
	response.Token = token

	return response, nil
}

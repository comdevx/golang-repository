package service

import (
	"errors"
	"project/helper"
	logs "project/helper"
	"project/repository"
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

	err = helper.ComparePassword(user.Password, body.Password)
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

func (s authenService) ChangePassword(body PasswordForm) (*PasswordResponse, error) {

	user, err := s.userRepo.GetByUser(body.Username)
	if err != nil {
		logs.Error(err)
		return nil, ErrServerError()
	}

	err = helper.ComparePassword(user.Password, body.OldPassword)
	if err != nil {
		return nil, errors.New("Password is not matched")
	}

	body.NewPassword, err = helper.Password(body.NewPassword)
	if err != nil {
		return nil, err
	}

	data := repository.UpdatePassword{
		ID:       user.ID,
		Username: user.Username,
		Password: body.NewPassword,
	}

	err = s.userRepo.UpdatePassword(data)
	if err != nil {
		return nil, ErrServerError()
	}

	result := PasswordResponse{
		Message: "Changed password success",
	}

	return &result, nil

}

package service

import (
	"errors"
	logs "project/helper"
	"project/repository"
	"unsafe"

	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) userService {
	return userService{userRepo: userRepo}
}

func (s userService) GetUsers(page, limit int) (UserListResponse, error) {

	page--

	users, err := s.userRepo.GetAll(page*limit, limit)
	if err != nil {
		logs.Error(err)
		return UserListResponse{}, ErrServerError()
	}

	userResponses := UserListResponse{}
	for _, user := range users.List {
		userResponse := UserResponse{
			ID:       user.ID,
			Username: user.Username,
		}
		userResponses.List = append(userResponses.List, userResponse)
	}
	userResponses.Total = users.Total

	return userResponses, nil
}

func (s userService) GetUser(id int) (*UserResponse, error) {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		logs.Error(err)
		return nil, ErrNotFoundError("user not found")
	}

	UserResponse := (*UserResponse)(unsafe.Pointer(user))

	return UserResponse, nil
}

func (s userService) NewUser(body NewUserRequest) (*UserResponse, error) {

	checkUser, err := s.userRepo.GetByUser(body.Username)
	if err != nil {
		return nil, ErrServerError()
	}

	if checkUser.Username == body.Username {
		return nil, errors.New("Username is Exists")
	}

	encode, _ := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)

	user := repository.User{
		Username: body.Username,
		Password: string(encode),
	}

	newUser, err := s.userRepo.Create(user)
	if err != nil {
		logs.Error(err)
		return nil, ErrServerError()
	}

	UserResponse := (*UserResponse)(unsafe.Pointer(newUser))

	return UserResponse, nil
}

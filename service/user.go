package service

import (
	logs "project/helper"
	"project/repository"
	"unsafe"
)

type NewUserRequest struct {
	UserID    int    `gorm="primaryKey"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Verified  bool   `json:"verified"`
	Suspended bool   `json:"suspended"`
}

type UserResponse struct {
	UserID    int    `gorm="primaryKey"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Verified  bool   `json:"verified"`
	Suspended bool   `json:"suspended"`
}

type UserService interface {
	GetUsers() ([]UserResponse, error)
	GetUser(id int) (*UserResponse, error)
	NewUser(NewUserRequest) (*UserResponse, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) userService {
	return userService{userRepo: userRepo}
}

func (s userService) GetUsers() ([]UserResponse, error) {

	users, err := s.userRepo.GetAll()
	if err != nil {
		logs.Error(err)
		return nil, ErrServerError()
	}

	userResponses := []UserResponse{}
	for _, user := range users {
		userResponse := UserResponse{
			UserID:   user.UserID,
			Username: user.Username,
			Password: user.Password,
		}
		userResponses = append(userResponses, userResponse)
	}

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

	user := repository.User{
		Username: body.Username,
		Password: body.Password,
	}

	newUser, err := s.userRepo.Create(user)
	if err != nil {
		logs.Error(err)
		return nil, ErrServerError()
	}

	UserResponse := (*UserResponse)(unsafe.Pointer(newUser))

	return UserResponse, nil
}

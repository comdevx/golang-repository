package service

import (
	logs "project/helper"
	"project/repository"
	"unsafe"
)

type NewUserRequest struct {
	ID        int    `gorm="primaryKey"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Verified  bool   `json:"verified"`
	Suspended bool   `json:"suspended"`
}

type UserResponse struct {
	ID        int    `gorm="primaryKey"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Verified  bool   `json:"verified"`
	Suspended bool   `json:"suspended"`
}

type UserListResponse struct {
	List  []UserResponse `json:"list"`
	Total int            `json:"total"`
}

type UserService interface {
	GetUsers(page, limit int) (UserListResponse, error)
	GetUser(id int) (*UserResponse, error)
	NewUser(NewUserRequest) (*UserResponse, error)
}

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
			Password: user.Password,
		}
		userResponses.List = append(userResponses.List, userResponse)
	}

	// userResponses.List = users.List
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

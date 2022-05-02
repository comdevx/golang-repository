package service

import (
	"errors"
	"project/helper"
	logs "project/helper"
	"project/repository"
	"strconv"
	"unsafe"
)

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) userService {
	return userService{userRepo: userRepo}
}

func (s userService) GetUsers(page, limit string) (UserListResponse, error) {

	pageToInt := 0
	limitToInt := 10

	if page != "" || limit != "" {
		pageToInt, _ = strconv.Atoi(page)
		limitToInt, _ = strconv.Atoi(limit)

		if pageToInt < 1 {
			return UserListResponse{}, errors.New("Page value less than 1")
		}

		if limitToInt < 1 || limitToInt > 100 {
			return UserListResponse{}, errors.New("Limit values less than 1 or greater than 100")
		}

		pageToInt--
	}

	users, err := s.userRepo.GetAll(pageToInt*limitToInt, limitToInt)
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

	encode, err := helper.Password(body.Password)
	if err != nil {
		logs.Error(err)
		return nil, ErrServerError()
	}

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

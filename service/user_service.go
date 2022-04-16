package service

import (
	"bank/errs"
	"bank/logs"
	"bank/repository"
	"unsafe"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userService struct {
	custRepo repository.UserRepository
}

func NewUserService(custRepo repository.UserRepository) userService {
	return userService{custRepo: custRepo}
}

func (s userService) GetUsers() ([]UserResponse, error) {

	users, err := s.custRepo.GetAll()
	if err != nil {
		logs.Error(err)
		return nil, errs.NewNotFoundError("user not found")
	}

	custResponses := []UserResponse{}
	for _, user := range users {
		custResponse := UserResponse{
			UserID:   user.UserID,
			Username: user.Username,
			Password: user.Password,
			Profile:  Profile(user.Profile),
		}
		custResponses = append(custResponses, custResponse)
	}

	return custResponses, nil
}

func (s userService) GetUser(id string) (*UserResponse, error) {

	user, err := s.custRepo.GetByID(id)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewNotFoundError("user not found")
	}

	UserResponse := (*UserResponse)(unsafe.Pointer(user))

	return UserResponse, nil
}

func (s userService) NewUser(body NewUserRequest) (*UserResponse, error) {

	if len(body.Username) < 4 {
		return nil, errs.NewValidationError("character at least 4")
	}

	if len(body.Password) < 6 {
		return nil, errs.NewValidationError("character at least 6")
	}

	user := repository.User{
		UserID:   primitive.NewObjectID(),
		Username: body.Username,
		Password: body.Password,
	}

	newUser, err := s.custRepo.Create(user)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewServerError()
	}

	UserResponse := (*UserResponse)(unsafe.Pointer(newUser))

	return UserResponse, nil
}

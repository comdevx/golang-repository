package service

// import (
// 	logs "project/helper"
// 	errs "project/helper/errs"
// 	"project/repository"
// 	"unsafe"

// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

// type userService struct {
// 	userRepo repository.UserRepository
// }

// func NewUserService(userRepo repository.UserRepository) userService {
// 	return userService{userRepo: userRepo}
// }

// func (s userService) GetUsers() ([]UserResponse, error) {

// 	users, err := s.userRepo.GetAll()
// 	if err != nil {
// 		logs.Error(err)
// 		return nil, errs.NewNotFoundError("user not found")
// 	}

// 	userResponses := []UserResponse{}
// 	for _, user := range users {
// 		userResponse := UserResponse{
// 			ID:   user.ID,
// 			Username: user.Username,
// 			Password: user.Password,
// 			Profile:  Profile(user.Profile),
// 		}
// 		userResponses = append(userResponses, userResponse)
// 	}

// 	return userResponses, nil
// }

// func (s userService) GetUser(id string) (*UserResponse, error) {

// 	user, err := s.userRepo.GetByID(id)
// 	if err != nil {
// 		logs.Error(err)
// 		return nil, errs.NewNotFoundError("user not found")
// 	}

// 	UserResponse := (*UserResponse)(unsafe.Pointer(user))

// 	return UserResponse, nil
// }

// func (s userService) NewUser(body NewUserRequest) (*UserResponse, error) {

// 	if len(body.Username) < 4 {
// 		return nil, errs.NewValidationError("character at least 4")
// 	}

// 	if len(body.Password) < 6 {
// 		return nil, errs.NewValidationError("character at least 6")
// 	}

// 	user := repository.User{
// 		ID:   primitive.NewObjectID(),
// 		Username: body.Username,
// 		Password: body.Password,
// 	}

// 	newUser, err := s.userRepo.Create(user)
// 	if err != nil {
// 		logs.Error(err)
// 		return nil, errs.NewServerError()
// 	}

// 	UserResponse := (*UserResponse)(unsafe.Pointer(newUser))

// 	return UserResponse, nil
// }

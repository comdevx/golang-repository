package service

import "go.mongodb.org/mongo-driver/bson/primitive"

type NewUserRequest struct {
	UserID    primitive.ObjectID `json:"id"`
	Username  string             `json:"username"`
	Password  string             `json:"password"`
	Verified  bool               `json:"verified"`
	Suspended bool               `json:"suspended"`
	Profile   Profile
}

type UserResponse struct {
	UserID    primitive.ObjectID `json:"id"`
	Username  string             `json:"username"`
	Password  string             `json:"password"`
	Verified  bool               `json:"verified"`
	Suspended bool               `json:"suspended"`
	Profile   Profile
}

type Profile struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       string `json:"age"`
}

type UserService interface {
	GetUsers() ([]UserResponse, error)
	GetUser(id string) (*UserResponse, error)
	NewUser(NewUserRequest) (*UserResponse, error)
}

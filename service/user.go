package service

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserResponse struct {
	UserID   primitive.ObjectID `json:"id"`
	Username string             `json:"username"`
	Password string             `json:"password"`
}

type NewUserRequest struct {
	UserID   primitive.ObjectID `json:"id"`
	Username string             `json:"username"`
	Password string             `json:"password"`
}

type UserService interface {
	GetCustomers() ([]UserResponse, error)
	NewUser()
}

package repository

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	UserID   primitive.ObjectID `bson:"_id" json:"id"`
	Username string             `json:"username"`
	Password string             `json:"password"`
}

type UserRepository interface {
	GetAll() ([]User, error)
	Create(id string) (*User, error)
}

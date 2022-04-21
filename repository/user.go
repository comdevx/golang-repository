package repository

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	UserID    primitive.ObjectID `bson:"_id" json:"id"`
	Username  string             `json:"username"`
	Password  string             `json:"password"`
	Verified  bool               `json:"verified"`
	Suspended bool               `json:"suspended"`
}

type UserRepository interface {
	GetAll() ([]User, error)
	GetByID(id string) (*User, error)
	Create(user User) (*User, error)
}

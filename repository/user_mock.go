package repository

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userRepositoryMock struct {
	users []User
}

func NewUserRepositoryMock() userRepositoryMock {
	users := []User{
		{UserID: primitive.NewObjectID(), Username: "test1", Password: "pass1", Verified: false, Suspended: false},
		{UserID: primitive.NewObjectID(), Username: "test2", Password: "pass2", Verified: false, Suspended: false},
		{UserID: primitive.NewObjectID(), Username: "test3", Password: "pass3", Verified: false, Suspended: false},
	}

	return userRepositoryMock{users: users}
}

func (r userRepositoryMock) GetAll() ([]User, error) {
	return r.users, nil
}

func (r userRepositoryMock) GetByID(id string) (*User, error) {

	convID, _ := primitive.ObjectIDFromHex(id)
	for _, user := range r.users {
		if user.UserID == convID {
			return &user, nil
		}
	}

	return nil, errors.New("user not found")
}

func (r userRepositoryMock) Create(user User) (*User, error) {

	user.UserID = primitive.NewObjectID()

	return &user, nil
}

package repository

import (
	"errors"
)

type userRepositoryMock struct {
	users []User
}

func NewUserRepositoryMock() userRepositoryMock {
	users := []User{
		{UserID: 1, Username: "test1", Password: "pass1", Verified: false, Suspended: false},
		{UserID: 2, Username: "test2", Password: "pass2", Verified: false, Suspended: false},
		{UserID: 3, Username: "test3", Password: "pass3", Verified: false, Suspended: false},
	}

	return userRepositoryMock{users: users}
}

func (r *userRepositoryMock) GetAll() ([]User, error) {
	return r.users, nil
}

func (r *userRepositoryMock) GetByID(id int) (*User, error) {

	for _, user := range r.users {
		if user.UserID == id {
			return &user, nil
		}
	}

	return nil, errors.New("user not found")
}

func (r *userRepositoryMock) Create(user User) (*User, error) {
	return &user, nil
}

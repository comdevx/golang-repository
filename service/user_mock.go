package service

import (
	"github.com/stretchr/testify/mock"
)

type UserServiceMock struct {
	mock.Mock
}

// func NewUserServiceMock() *userServiceMock {
// 	return &userServiceMock{}
// }

func (m *UserServiceMock) Login(body Body) (*UserResponse, error) {
	args := m.Called(body)
	return args.Get(0).(*UserResponse), args.Error(1)
}

func (m *UserServiceMock) GetUsers(page, limit string) (UserListResponse, error) {
	args := m.Called(page, limit)
	return args.Get(0).(UserListResponse), args.Error(1)
}

func (m *UserServiceMock) GetUser(id int) (*UserResponse, error) {
	args := m.Called(id)
	return args.Get(0).(*UserResponse), args.Error(1)
}

func (m *UserServiceMock) NewUser(body NewUserRequest) (*UserResponse, error) {
	args := m.Called(body)
	return args.Get(0).(*UserResponse), args.Error(1)
}

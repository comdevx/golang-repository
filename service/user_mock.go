package service

import (
	"github.com/stretchr/testify/mock"
)

type userServiceMock struct {
	mock.Mock
}

func NewUserServiceMock() *userServiceMock {
	return &userServiceMock{}
}

func (m *userServiceMock) GetUsers() ([]UserResponse, error) {
	args := m.Called()
	return args.Get(0).([]UserResponse), args.Error(1)
}

func (m *userServiceMock) GetUser(id string) (*UserResponse, error) {
	args := m.Called(id)
	return args.Get(0).(*UserResponse), args.Error(1)
}

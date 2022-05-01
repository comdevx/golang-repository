package repository

import "time"

type User struct {
	ID        int `gorm:"primaryKey"`
	Username  string
	Password  string
	Verified  bool
	Suspended bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Users struct {
	List  []User
	Total int
}

type Body struct {
	Username string
	Password string
}

type UpdatePassword struct {
	ID        int
	Username  string
	Password  string
	UpdatedAt time.Time
}

type UserRepository interface {
	GetAll(skip, limit int) (Users, error)
	GetByID(id int) (*User, error)
	GetByUser(string) (*User, error)
	Create(User) (*User, error)
	UpdatePassword(UpdatePassword) error
}

package repository

type User struct {
	ID        int `gorm:"primaryKey"`
	Username  string
	Password  string
	Verified  bool
	Suspended bool
}

type Users struct {
	List  []User
	Total int
}

type UserRepository interface {
	GetAll(skip, limit int) (Users, error)
	GetByID(id int) (*User, error)
	Create(user User) (*User, error)
}

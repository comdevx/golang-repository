package repository

type User struct {
	UserID    int    `json:"id" gorm:"primaryKey"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Verified  bool   `json:"verified"`
	Suspended bool   `json:"suspended"`
}

type UserRepository interface {
	GetAll() ([]User, error)
	GetByID(id int) (*User, error)
	Create(user User) (*User, error)
}

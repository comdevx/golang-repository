package repository

type User struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Verified  bool   `json:"verified"`
	Suspended bool   `json:"suspended"`
}

type Users struct {
	List  []User `json:"list"`
	Total int    `json:"total"`
}

type UserRepository interface {
	GetAll(skip, limit int) (Users, error)
	GetByID(id int) (*User, error)
	Create(user User) (*User, error)
}

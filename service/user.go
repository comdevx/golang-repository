package service

type NewUserRequest struct {
	ID        int    `gorm="primaryKey"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Verified  bool   `json:"verified"`
	Suspended bool   `json:"suspended"`
}

type UserResponse struct {
	ID        int    `gorm="primaryKey"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Verified  bool   `json:"verified"`
	Suspended bool   `json:"suspended"`
}

type UserListResponse struct {
	List  []UserResponse `json:"list"`
	Total int            `json:"total"`
}

type UserService interface {
	GetUsers(page, limit int) (UserListResponse, error)
	GetUser(id int) (*UserResponse, error)
	NewUser(NewUserRequest) (*UserResponse, error)
}

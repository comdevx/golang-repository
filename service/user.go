package service

type Body struct {
	Username string `json:"username" biding:"min=4"`
	Password string `json:"password" biding:"min=6"`
}

type NewUserRequest struct {
	ID        int    `json:"id"`
	Username  string `json:"username" binding:"required,min=4"`
	Password  string `json:"password" binding:"required,min=6"`
	Verified  bool   `json:"verified"`
	Suspended bool   `json:"suspended"`
}

type UserResponse struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"-"`
	Verified  bool   `json:"verified"`
	Suspended bool   `json:"suspended"`
}

type UserListResponse struct {
	List  []UserResponse `json:"list"`
	Total int            `json:"total"`
}

type UserService interface {
	GetUsers(page, limit string) (UserListResponse, error)
	GetUser(id int) (*UserResponse, error)
	NewUser(NewUserRequest) (*UserResponse, error)
}

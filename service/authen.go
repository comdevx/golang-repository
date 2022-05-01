package service

type AuthenBody struct {
	Username string `json:"username" biding:"min=4"`
	Password string `json:"password" biding:"min=6"`
}

type AuthenResponse struct {
	Token string `json:"token"`
}

type PasswordForm struct {
	Username    string `json:"username" biding:"min=4"`
	OldPassword string `json:"old_password" biding:"min=6"`
	NewPassword string `json:"new_password" biding:"min=6"`
}

type PasswordResponse struct {
	Message string `json:"message"`
}

type AuthenService interface {
	Login(AuthenBody) (*AuthenResponse, error)
	ChangePassword(PasswordForm) (*PasswordResponse, error)
}

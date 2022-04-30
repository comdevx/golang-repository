package service

type AuthenBody struct {
	Username string `json:"username" biding:"min=4"`
	Password string `json:"password" biding:"min=6"`
}

type AuthenResponse struct {
	Token string `json:"token"`
}

type AuthenService interface {
	Login(AuthenBody) (*AuthenResponse, error)
}

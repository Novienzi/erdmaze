package request

import "erdmaze/businesses/users"

type Users struct {
	Fullname string `json:"fullname"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (req *Users) ToDomain() *users.Domain {
	return &users.Domain{
		Fullname: req.Fullname,
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}
}

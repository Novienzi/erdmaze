package response

import (
	"erdmaze/businesses/users"
	"time"
)

type Users struct {
	Id        int       `json:"id"`
	Fullname  string    `json:"fullname"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomain(domain users.Domain) Users {
	return Users{
		Id:        domain.Id,
		Fullname:  domain.Fullname,
		Username:  domain.Username,
		Email:     domain.Email,
		Token:     domain.Token,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

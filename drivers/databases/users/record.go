package users

import (
	"erdmaze/businesses/users"
	"time"
)

type Users struct {
	ID        int
	Fullname  string
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (rec *Users) toDomain() users.Domain {
	return users.Domain{
		Id:        rec.ID,
		Fullname:  rec.Fullname,
		Username:  rec.Username,
		Email:     rec.Email,
		Password:  rec.Password,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func fromDomain(userDomain users.Domain) *Users {
	return &Users{
		ID:        userDomain.Id,
		Fullname:  userDomain.Fullname,
		Username:  userDomain.Username,
		Email:     userDomain.Email,
		Password:  userDomain.Password,
		CreatedAt: userDomain.CreatedAt,
		UpdatedAt: userDomain.UpdatedAt,
	}
}

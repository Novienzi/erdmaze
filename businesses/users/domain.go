package users

import (
	"context"
	"time"
)

type Domain struct {
	Id        int
	Fullname  string
	Username  string
	Email     string
	Password  string
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface {
	Login(ctx context.Context, username, password string) (string, error)
	Store(ctx context.Context, data *Domain) error
	GetByID(ctx context.Context, newsId int) (Domain, error)
}

type Repository interface {
	GetByUsername(ctx context.Context, username string) (Domain, error)
	Store(ctx context.Context, data *Domain) error
	GetByID(ctx context.Context, newsId int) (Domain, error)
}

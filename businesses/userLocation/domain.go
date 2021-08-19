package user_locations

import (
	"context"
	"time"
)

type Domain struct {
	ID        int
	UserID    int
	Region    string
	City      string
	CreatedAt time.Time
	UpdatedAt time.Time
	IPStat    string
}

type Usecase interface {
	GetByUserID(ctx context.Context, userID int) (Domain, error)
	Store(ctx context.Context, ip string, userLocationDomain *Domain) (Domain, error)
}

type Repository interface {
	GetByUserID(ctx context.Context, userID int) (Domain, error)
	GetByID(ctx context.Context, ID int) (Domain, error)
	Store(ctx context.Context, userLocationDomain *Domain) (Domain, error)
}

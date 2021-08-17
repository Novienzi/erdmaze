package bookings

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID                  int
	UserID              int
	Username            string
	TourismPackageID    int
	TourismsPackageName string
	StartTime           time.Time
	EndTime             time.Time
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           gorm.DeletedAt
}

type Usecase interface {
	Fetch(ctx context.Context, page, perpage int) ([]Domain, int, int, error)
	Store(ctx context.Context, bookingsDomain *Domain) (Domain, error)
	GetByUserID(ctx context.Context, userID int) ([]Domain, error)
	GetByID(ctx context.Context, locationId int) (Domain, error)
}

type Repository interface {
	Fetch(ctx context.Context, page, perpage int) ([]Domain, int, error)
	GetByID(ctx context.Context, bookingsId int) (Domain, error)
	Store(ctx context.Context, bookingsDomain *Domain) (Domain, error)
	GetByUserID(ctx context.Context, userID int) ([]Domain, error)
}

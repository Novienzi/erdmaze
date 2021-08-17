package bookings

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID               int
	UserID           int
	User             string
	TourismPackageID int
	TourismsPackage  string
	StartTime        time.Time
	EndTime          time.Time
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt
}

type Usecase interface {
	GetByUserID(ctx context.Context, userID int) ([]Domain, error)
	GetByID(ctx context.Context, ID int) (Domain, error)
	Store(ctx context.Context, bookingDomain *Domain) (Domain, error)
	Delete(ctx context.Context, bookingDomain *Domain) (*Domain, error)
}

type Repository interface {
	GetByUserID(ctx context.Context, userID int) ([]Domain, error)
	GetByID(ctx context.Context, ID int) (Domain, error)
	Store(ctx context.Context, bookingDomain *Domain) (Domain, error)
	Delete(ctx context.Context, bookingDomain *Domain) (Domain, error)
}

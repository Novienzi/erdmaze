package tourismpackages

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID           int
	Name         string
	Description  string
	TotalPrice   string
	TotalTime    string
	LocationID   int
	LocationName string
	ActivityID   int
	ActivityName string
	Address      string
	AddressUrl   string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
}

type Usecase interface {
	Fetch(ctx context.Context, page, perpage int) ([]Domain, int, error)
	Store(ctx context.Context, tourismPackagesDomain *Domain) (Domain, error)
	GetAll(ctx context.Context, tourismName string, locationName string, activityName string) ([]Domain, error)
	GetByID(ctx context.Context, locationId int) (Domain, error)
}

type Repository interface {
	Fetch(ctx context.Context, page, perpage int) ([]Domain, int, error)
	GetByID(ctx context.Context, tourismPackagesId int) (Domain, error)
	GetByName(ctx context.Context, tourismPackagesName string) (Domain, error)
	Store(ctx context.Context, tourismPackagesDomain *Domain) (Domain, error)
	GetAll(ctx context.Context, tourismName string, locationName string, activityName string) ([]Domain, error)
}

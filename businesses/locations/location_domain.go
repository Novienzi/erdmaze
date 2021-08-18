package locations

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type Usecase interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetByID(ctx context.Context, locationId int) (Domain, error)
	Store(ctx context.Context, locationDomain *Domain) (Domain, error)
}

type Repository interface {
	Find(ctx context.Context) ([]Domain, error)
	GetByID(ctx context.Context, locationId int) (Domain, error)
	GetByName(ctx context.Context, locationName string) (Domain, error)
	Store(ctx context.Context, locationDomain *Domain) (Domain, error)
}

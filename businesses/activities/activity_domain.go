package activities

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
	Fetch(ctx context.Context, page, perpage int) ([]Domain, int, error)
	GetAll(ctx context.Context) ([]Domain, error)
	GetByID(ctx context.Context, activityId int) (Domain, error)
	GetByName(ctx context.Context, activityName string) (Domain, error)
	Store(ctx context.Context, activityDomain *Domain) (Domain, error)
	Update(ctx context.Context, activityDomain *Domain) (*Domain, error)
	Delete(ctx context.Context, ActivitiesDomain *Domain) (*Domain, error)
}

type Repository interface {
	Fetch(ctx context.Context, page, perpage int) ([]Domain, int, error)
	Find(ctx context.Context) ([]Domain, error)
	GetByID(ctx context.Context, activityId int) (Domain, error)
	GetByName(ctx context.Context, activityName string) (Domain, error)
	Store(ctx context.Context, ActivitiesDomain *Domain) (Domain, error)
	Update(ctx context.Context, ActivitiesDomain *Domain) (Domain, error)
	Delete(ctx context.Context, ActivitiesDomain *Domain) (Domain, error)
}

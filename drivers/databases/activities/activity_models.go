package activities

import (
	activitiesUsecase "erdmaze/businesses/activities"
	"time"

	"gorm.io/gorm"
)

type Activities struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (rec *Activities) toDomain() activitiesUsecase.Domain {
	return activitiesUsecase.Domain{
		ID:        rec.ID,
		Name:      rec.Name,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
		DeletedAt: rec.DeletedAt,
	}
}

func fromDomain(ActivitiesDomain *activitiesUsecase.Domain) *Activities {
	return &Activities{
		ID:        ActivitiesDomain.ID,
		Name:      ActivitiesDomain.Name,
		CreatedAt: ActivitiesDomain.CreatedAt,
		UpdatedAt: ActivitiesDomain.UpdatedAt,
		DeletedAt: ActivitiesDomain.DeletedAt,
	}
}

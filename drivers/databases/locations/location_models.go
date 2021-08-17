package locations

import (
	locationsUsecase "erdmaze/businesses/locations"
	"time"

	"gorm.io/gorm"
)

type Locations struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (rec *Locations) toDomain() locationsUsecase.Domain {
	return locationsUsecase.Domain{
		ID:        rec.ID,
		Name:      rec.Name,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
		DeletedAt: rec.DeletedAt,
	}
}

func fromDomain(LocationsDomain *locationsUsecase.Domain) *Locations {
	return &Locations{
		ID:        LocationsDomain.ID,
		Name:      LocationsDomain.Name,
		CreatedAt: LocationsDomain.CreatedAt,
		UpdatedAt: LocationsDomain.UpdatedAt,
		DeletedAt: LocationsDomain.DeletedAt,
	}
}

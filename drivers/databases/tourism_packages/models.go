package tourismpackages

import (
	tourismpackagesUsecase "erdmaze/businesses/tourism_packages"
	"time"

	"gorm.io/gorm"
)

type TourismPackages struct {
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

func (rec *TourismPackages) toDomain() tourismpackagesUsecase.Domain {
	return tourismpackagesUsecase.Domain{
		ID:           rec.ID,
		Name:         rec.Name,
		Description:  rec.Description,
		TotalPrice:   rec.TotalPrice,
		TotalTime:    rec.TotalTime,
		LocationID:   rec.LocationID,
		LocationName: rec.LocationName,
		ActivityID:   rec.ActivityID,
		ActivityName: rec.ActivityName,
		Address:      rec.Address,
		AddressUrl:   rec.AddressUrl,
		CreatedAt:    rec.CreatedAt,
		UpdatedAt:    rec.UpdatedAt,
		DeletedAt:    rec.DeletedAt,
	}
}

func fromDomain(tourismDomain *tourismpackagesUsecase.Domain) *TourismPackages {
	return &TourismPackages{
		ID:          tourismDomain.ID,
		Name:        tourismDomain.Name,
		Description: tourismDomain.Description,
		TotalPrice:  tourismDomain.TotalTime,
		TotalTime:   tourismDomain.TotalTime,
		LocationID:  tourismDomain.LocationID,
		ActivityID:  tourismDomain.ActivityID,
		Address:     tourismDomain.Address,
		AddressUrl:  tourismDomain.AddressUrl,
		CreatedAt:   tourismDomain.CreatedAt,
		UpdatedAt:   tourismDomain.UpdatedAt,
		DeletedAt:   tourismDomain.DeletedAt,
	}
}

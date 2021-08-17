package response

import (
	tourismpackages "erdmaze/businesses/tourism_packages"
	"time"

	"gorm.io/gorm"
)

type TourismPackages struct {
	Id           int            `json:"id"`
	Name         string         `json:"name"`
	Description  string         `json:"description"`
	TotalPrice   string         `json:"total_price"`
	TotalTime    string         `json:"total_time"`
	LocationID   int            `json:"location_id"`
	LocationName string         `json:"location_name"`
	ActivityID   int            `json:"activity_id"`
	ActivityName string         `json:"activity_name"`
	Address      string         `json:"address"`
	AddressUrl   string         `json:"address_url"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
}

func FromDomain(domain tourismpackages.Domain) TourismPackages {
	return TourismPackages{
		Id:           domain.ID,
		Name:         domain.Name,
		Description:  domain.Description,
		TotalPrice:   domain.TotalPrice,
		TotalTime:    domain.TotalTime,
		LocationID:   domain.LocationID,
		LocationName: domain.LocationName,
		ActivityID:   domain.ActivityID,
		ActivityName: domain.ActivityName,
		Address:      domain.Address,
		AddressUrl:   domain.AddressUrl,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
		DeletedAt:    domain.DeletedAt,
	}
}

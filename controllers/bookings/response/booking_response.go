package response

import (
	bookings "erdmaze/businesses/bookings"
	tourismpackages "erdmaze/drivers/databases/tourism_packages"
	"erdmaze/drivers/databases/users"
	"time"

	"gorm.io/gorm"
)

type Bookings struct {
	Id               int                             `json:"id"`
	UserID           int                             `json:"user_id"`
	Users            users.Users                     `json:"users"`
	TourismPackageID int                             `json:"tourism_package_id"`
	TourismsName     tourismpackages.TourismPackages `json:"tourism_packages"`
	StartTime        time.Time                       `json:"Start_time"`
	EndTime          time.Time                       `json:"End_time"`
	CreatedAt        time.Time                       `json:"created_at"`
	UpdatedAt        time.Time                       `json:"updated_at"`
	DeletedAt        gorm.DeletedAt                  `json:"deleted_at"`
}

func FromDomain(domain bookings.Domain) Bookings {
	return Bookings{
		Id:               domain.ID,
		UserID:           domain.UserID,
		TourismPackageID: domain.TourismPackageID,
		StartTime:        domain.StartTime,
		EndTime:          domain.EndTime,
		CreatedAt:        domain.CreatedAt,
		UpdatedAt:        domain.UpdatedAt,
		DeletedAt:        domain.DeletedAt,
	}
}

type Pagination struct {
	CurrentPage int `json:"current_page"`
	LastPage    int `json:"last_page"`
	Total       int `json:"total"`
	PerPage     int `json:"per_page"`
}

package response

import (
	bookings "erdmaze/businesses/bookings"
	"time"

	"gorm.io/gorm"
)

type Bookings struct {
	Id               int            `json:"id"`
	UserID           int            `json:"user_id"`
	User             string         `json:"user"`
	TourismPackageID int            `json:"tourism_package_id"`
	TourismPackage   string         `json:"tourism_package`
	StartTime        time.Time      `json:"Start_time"`
	EndTime          time.Time      `json:"End_time"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `json:"deleted_at"`
}

func FromDomain(domain bookings.Domain) Bookings {
	return Bookings{
		Id:               domain.ID,
		UserID:           domain.UserID,
		User:             domain.User,
		TourismPackageID: domain.TourismPackageID,
		TourismPackage:   domain.TourismsPackage,
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

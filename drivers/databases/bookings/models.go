package bookings

import (
	bookingsUsecase "erdmaze/businesses/bookings"
	"time"

	"gorm.io/gorm"
)

type Bookings struct {
	ID               int
	UserID           int
	User             string `json:"users" gorm:"foreignKey:UserID;references:ID"`
	TourismPackageID int    `json:"torism_package_id"`
	TourismPackage   string `json:"tourism_packages" gorm:"foreignKey:TourismPackagesID;references:ID"`
	StartTime        time.Time
	EndTime          time.Time
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt
}

func (rec *Bookings) toDomain() bookingsUsecase.Domain {
	return bookingsUsecase.Domain{
		ID:               rec.ID,
		UserID:           rec.UserID,
		User:             rec.User,
		TourismPackageID: rec.TourismPackageID,
		TourismsPackage:  rec.TourismPackage,
		StartTime:        rec.StartTime,
		EndTime:          rec.EndTime,
		CreatedAt:        rec.CreatedAt,
		UpdatedAt:        rec.UpdatedAt,
		DeletedAt:        rec.DeletedAt,
	}
}

func fromDomain(bookingDomain *bookingsUsecase.Domain) *Bookings {
	return &Bookings{
		ID:               bookingDomain.ID,
		UserID:           bookingDomain.UserID,
		User:             bookingDomain.User,
		TourismPackageID: bookingDomain.TourismPackageID,
		TourismPackage:   bookingDomain.TourismsPackage,
		StartTime:        bookingDomain.StartTime,
		EndTime:          bookingDomain.EndTime,
		CreatedAt:        bookingDomain.CreatedAt,
		UpdatedAt:        bookingDomain.UpdatedAt,
		DeletedAt:        bookingDomain.DeletedAt,
	}
}

package bookings

import (
	bookingsUsecase "erdmaze/businesses/bookings"
	"time"

	"gorm.io/gorm"
)

type Bookings struct {
	ID               int
	UserID           int
	TourismPackageID int
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
		TourismPackageID: rec.TourismPackageID,
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
		TourismPackageID: bookingDomain.TourismPackageID,
		StartTime:        bookingDomain.StartTime,
		EndTime:          bookingDomain.EndTime,
		CreatedAt:        bookingDomain.CreatedAt,
		UpdatedAt:        bookingDomain.UpdatedAt,
		DeletedAt:        bookingDomain.DeletedAt,
	}
}

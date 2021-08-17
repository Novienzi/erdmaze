package request

import (
	bookings "erdmaze/businesses/bookings"
	"time"
)

type Bookings struct {
	Id               int       `json:"id"`
	UserID           int       `json:"user_id"`
	TourismPackageID int       `json:"tourism_package_id"`
	StartTime        time.Time `json:"Start_time"`
	EndTime          time.Time `json:"End_time"`
}

func (req *Bookings) ToDomain() *bookings.Domain {
	return &bookings.Domain{
		UserID:           req.UserID,
		TourismPackageID: req.TourismPackageID,
		StartTime:        req.StartTime,
		EndTime:          req.EndTime,
	}
}

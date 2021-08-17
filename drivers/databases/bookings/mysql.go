package bookings

import (
	"context"
	bookings "erdmaze/businesses/bookings"

	"gorm.io/gorm"
)

type mysqlBookingsRepository struct {
	Conn *gorm.DB
}

func NewMySQLBookingsRepository(conn *gorm.DB) bookings.Repository {
	return &mysqlBookingsRepository{
		Conn: conn,
	}
}

func (nr *mysqlBookingsRepository) Fetch(ctx context.Context, page, perpage int) ([]bookings.Domain, int, error) {
	rec := []Bookings{}

	offset := (page - 1) * perpage

	err := nr.Conn.Select("tourism_packages.*, activities.name as activity_name , locations.name as location_name").
		Joins("JOIN activities on activities.id = tourism_packages.activity_id").
		Joins("JOIN locations on locations.id = tourism_packages.location_id").
		Find(&rec).Offset(offset).Limit(perpage).Error

	if err != nil {
		return []bookings.Domain{}, 0, err
	}

	var totalData int64
	err = nr.Conn.Count(&totalData).Error
	if err != nil {
		return []bookings.Domain{}, 0, err
	}

	var domainTourisms []bookings.Domain
	for _, value := range rec {
		domainTourisms = append(domainTourisms, value.toDomain())
	}
	return domainTourisms, int(totalData), nil
}

func (nr *mysqlBookingsRepository) GetByID(ctx context.Context, tourismId int) (bookings.Domain, error) {
	rec := Bookings{}
	err := nr.Conn.Select("bookings.*, ").
		Joins("JOIN activities on activities.id = tourism_packages.activity_id").
		Joins("JOIN locations on locations.id = tourism_packages.location_id").
		Where("tourism_packages.id = ?", tourismId).First(&rec).Error

	// err := nr.Conn.Joins("Activities").
	if err != nil {
		return bookings.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (cr *mysqlBookingsRepository) GetByUserID(ctx context.Context, userID int) ([]bookings.Domain, error) {
	rec := []Bookings{}

	err := cr.Conn.Joins("Users").Joins("tourism_packages").Where("bookings.user_id = ?", userID).Find(&rec).Error
	if err != nil {
		return []bookings.Domain{}, err
	}

	favouriteDomain := []bookings.Domain{}
	for _, value := range rec {
		favouriteDomain = append(favouriteDomain, value.toDomain())
	}

	return favouriteDomain, nil
}

func (nr *mysqlBookingsRepository) Store(ctx context.Context, bookingsDomain *bookings.Domain) (bookings.Domain, error) {
	rec := fromDomain(bookingsDomain)

	result := nr.Conn.Create(&rec)
	if result.Error != nil {
		return bookings.Domain{}, result.Error
	}

	err := nr.Conn.Preload("Users").First(&rec, rec.UserID).Error
	if err != nil {
		return bookings.Domain{}, result.Error
	}

	error := nr.Conn.Preload("TourismPackages").First(&rec, rec.TourismPackageID).Error
	if error != nil {
		return bookings.Domain{}, result.Error
	}

	return rec.toDomain(), nil
}

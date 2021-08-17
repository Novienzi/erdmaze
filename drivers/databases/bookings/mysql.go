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
func (repo *mysqlBookingsRepository) GetByUserID(ctx context.Context, userID int) ([]bookings.Domain, error) {
	rec := []Bookings{}

	err := repo.Conn.Select("bookings.*, activities.name as activity_name , locations.name as location_name").
		Joins("Join users on bookings.user_id = users.id").
		Joins("tourism_packages on bookings.tourism_package_id = tourism_packages.id").
		Joins("JOIN activities on activities.id = tourism_packages.activity_id").
		Joins("JOIN locations on locations.id = tourism_packages.location_id").
		Where("bookings.id = ?", userID).Find(&rec).Error

	if err != nil {
		return []bookings.Domain{}, err
	}

	bookingDomain := []bookings.Domain{}
	for _, value := range rec {
		bookingDomain = append(bookingDomain, value.toDomain())
	}

	return bookingDomain, nil
}

func (repo *mysqlBookingsRepository) GetByID(ctx context.Context, ID int) (bookings.Domain, error) {
	rec := Bookings{}
	err := repo.Conn.Joins("Users").Joins("TourismPackage").Where("bookings.id = ?", ID).First(&rec).Error
	if err != nil {
		return bookings.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (repo *mysqlBookingsRepository) Store(ctx context.Context, bookingsDomain *bookings.Domain) (bookings.Domain, error) {
	rec := fromDomain(bookingsDomain)

	result := repo.Conn.Create(&rec)
	if result.Error != nil {
		return bookings.Domain{}, result.Error
	}

	return rec.toDomain(), nil
}

func (repo *mysqlBookingsRepository) Delete(ctx context.Context, bookingDomain *bookings.Domain) (bookings.Domain, error) {
	rec := fromDomain(bookingDomain)

	result := repo.Conn.Delete(rec)

	if result.Error != nil {
		return bookings.Domain{}, result.Error
	}

	return rec.toDomain(), nil
}

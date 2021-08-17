package locations

import (
	"context"
	"erdmaze/businesses/locations"

	"gorm.io/gorm"
)

type mysqlLocationsRepository struct {
	DB *gorm.DB
}

func NewMySQLLocationRepository(conn *gorm.DB) locations.Repository {
	return &mysqlLocationsRepository{
		DB: conn,
	}
}

func (cr *mysqlLocationsRepository) Fetch(ctx context.Context, page, perpage int) ([]locations.Domain, int, error) {
	rec := []Locations{}

	offset := (page - 1) * perpage
	err := cr.DB.Find(&rec).Offset(offset).Limit(perpage).Error
	if err != nil {
		return []locations.Domain{}, 0, err
	}

	var totalData int64
	err = cr.DB.Model(&rec).Count(&totalData).Error
	if err != nil {
		return []locations.Domain{}, 0, err
	}

	var domainLocation []locations.Domain
	for _, value := range rec {
		domainLocation = append(domainLocation, value.toDomain())
	}
	return domainLocation, int(totalData), nil
}

func (cr *mysqlLocationsRepository) Find(ctx context.Context) ([]locations.Domain, error) {
	rec := []Locations{}

	cr.DB.Find(&rec)
	locationDomain := []locations.Domain{}
	for _, value := range rec {
		locationDomain = append(locationDomain, value.toDomain())
	}

	return locationDomain, nil
}

func (cr *mysqlLocationsRepository) GetByID(ctx context.Context, userId int) (locations.Domain, error) {
	rec := Locations{}
	err := cr.DB.Where("id = ?", userId).First(&rec).Error
	if err != nil {
		return locations.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (nr *mysqlLocationsRepository) GetByName(ctx context.Context, LocationName string) (locations.Domain, error) {
	rec := Locations{}
	err := nr.DB.Where("name = ?", LocationName).First(&rec).Error
	if err != nil {
		return locations.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (nr *mysqlLocationsRepository) Store(ctx context.Context, LocationDomain *locations.Domain) (locations.Domain, error) {
	rec := fromDomain(LocationDomain)

	result := nr.DB.Select("Name", "CreatedAt").Create(&rec)
	if result.Error != nil {
		return locations.Domain{}, result.Error
	}

	return rec.toDomain(), nil
}

func (nr *mysqlLocationsRepository) Update(ctx context.Context, LocationDomain *locations.Domain) (locations.Domain, error) {
	rec := fromDomain(LocationDomain)

	result := nr.DB.Updates(rec)
	if result.Error != nil {
		return locations.Domain{}, result.Error
	}

	return rec.toDomain(), nil
}

func (nr *mysqlLocationsRepository) Delete(ctx context.Context, LocationDomain *locations.Domain) (locations.Domain, error) {
	rec := fromDomain(LocationDomain)

	result := nr.DB.Delete(rec)

	if result.Error != nil {
		return locations.Domain{}, result.Error
	}

	return rec.toDomain(), nil
}

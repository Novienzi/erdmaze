package tourismpackages

import (
	"context"
	tourismpackages "erdmaze/businesses/tourism_packages"

	"gorm.io/gorm"
)

type mysqlTourismPackagesRepository struct {
	Conn *gorm.DB
}

func NewMySQLTourismPackagesRepository(conn *gorm.DB) tourismpackages.Repository {
	return &mysqlTourismPackagesRepository{
		Conn: conn,
	}
}

func (nr *mysqlTourismPackagesRepository) Fetch(ctx context.Context, page, perpage int) ([]tourismpackages.Domain, int, error) {
	rec := []TourismPackages{}

	offset := (page - 1) * perpage

	err := nr.Conn.Select("tourism_packages.*, activities.name as activity_name , locations.name as location_name").
		Joins("JOIN activities on activities.id = tourism_packages.activity_id").
		Joins("JOIN locations on locations.id = tourism_packages.location_id").
		Find(&rec).Offset(offset).Limit(perpage).Error

	if err != nil {
		return []tourismpackages.Domain{}, 0, err
	}

	var totalData int64
	err = nr.Conn.Count(&totalData).Error
	if err != nil {
		return []tourismpackages.Domain{}, 0, err
	}

	var domainTourisms []tourismpackages.Domain
	for _, value := range rec {
		domainTourisms = append(domainTourisms, value.toDomain())
	}
	return domainTourisms, int(totalData), nil
}

func (nr *mysqlTourismPackagesRepository) GetByID(ctx context.Context, tourismId int) (tourismpackages.Domain, error) {
	rec := TourismPackages{}
	err := nr.Conn.Select("tourism_packages.*, activities.name as activity_name , locations.name as location_name").
		Joins("JOIN activities on activities.id = tourism_packages.activity_id").
		Joins("JOIN locations on locations.id = tourism_packages.location_id").
		Where("tourism_packages.id = ?", tourismId).First(&rec).Error

	// err := nr.Conn.Joins("Activities").
	if err != nil {
		return tourismpackages.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (cr *mysqlTourismPackagesRepository) GetAll(ctx context.Context, tourismName string, locationName string, activityName string) ([]tourismpackages.Domain, error) {
	rec := []TourismPackages{}

	if tourismName != " " {
		err := cr.Conn.Select("tourism_packages.*, activities.name as activity_name , locations.name as location_name").Debug().
			Joins("JOIN activities on activities.id = tourism_packages.activity_id").
			Joins("JOIN locations on locations.id = tourism_packages.location_id").
			Where("tourism_packages.name LIKE ?", "%"+tourismName+"%").Find(&rec).Error

		if err != nil {
			return []tourismpackages.Domain{}, err
		}
	}

	if locationName != " " {
		err := cr.Conn.Select("tourism_packages.*, activities.name as activity_name , locations.name as location_name").Debug().
			Joins("JOIN activities on activities.id = tourism_packages.activity_id").
			Joins("JOIN locations on locations.id = tourism_packages.location_id").
			Where("locations.name LIKE ?", "%"+locationName+"%").Find(&rec).Error

		if err != nil {
			return []tourismpackages.Domain{}, err
		}
	}

	if activityName != " " {
		err := cr.Conn.Select("tourism_packages.*, activities.name as activity_name , locations.name as location_name").Debug().
			Joins("JOIN activities on activities.id = tourism_packages.activity_id").
			Joins("JOIN locations on locations.id = tourism_packages.location_id").
			Where("activities.name LIKE ?", "%"+activityName+"%").Find(&rec).Error

		if err != nil {
			return []tourismpackages.Domain{}, err
		}
	}

	tourismDomain := []tourismpackages.Domain{}
	for _, value := range rec {
		tourismDomain = append(tourismDomain, value.toDomain())
	}

	return tourismDomain, nil
}

func (nr *mysqlTourismPackagesRepository) Store(ctx context.Context, tourismPackagesDomain *tourismpackages.Domain) (tourismpackages.Domain, error) {
	rec := fromDomain(tourismPackagesDomain)

	result := nr.Conn.Create(&rec)
	if result.Error != nil {
		return tourismpackages.Domain{}, result.Error
	}

	err := nr.Conn.Preload("Activities").First(&rec, rec.ID).Error
	if err != nil {
		return tourismpackages.Domain{}, result.Error
	}

	error := nr.Conn.Preload("Locations").First(&rec, rec.ID).Error
	if error != nil {
		return tourismpackages.Domain{}, result.Error
	}

	return rec.toDomain(), nil
}

func (nr *mysqlTourismPackagesRepository) GetByName(ctx context.Context, tourismName string) (tourismpackages.Domain, error) {
	rec := TourismPackages{}
	err := nr.Conn.Where("name = ?", tourismName).First(&rec).Error
	if err != nil {
		return tourismpackages.Domain{}, err
	}
	return rec.toDomain(), nil
}

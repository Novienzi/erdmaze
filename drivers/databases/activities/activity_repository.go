package activities

import (
	"context"
	"erdmaze/businesses/activities"

	"gorm.io/gorm"
)

type mysqlActivitiesRepository struct {
	DB *gorm.DB
}

func NewMySQLActivityRepository(conn *gorm.DB) activities.Repository {
	return &mysqlActivitiesRepository{
		DB: conn,
	}
}

func (cr *mysqlActivitiesRepository) Fetch(ctx context.Context, page, perpage int) ([]activities.Domain, int, error) {
	rec := []Activities{}

	offset := (page - 1) * perpage
	err := cr.DB.Find(&rec).Offset(offset).Limit(perpage).Error
	if err != nil {
		return []activities.Domain{}, 0, err
	}

	var totalData int64
	err = cr.DB.Model(&rec).Count(&totalData).Error
	if err != nil {
		return []activities.Domain{}, 0, err
	}

	var domainActivity []activities.Domain
	for _, value := range rec {
		domainActivity = append(domainActivity, value.toDomain())
	}
	return domainActivity, int(totalData), nil
}

func (cr *mysqlActivitiesRepository) Find(ctx context.Context) ([]activities.Domain, error) {
	rec := []Activities{}

	cr.DB.Find(&rec)
	ActivityDomain := []activities.Domain{}
	for _, value := range rec {
		ActivityDomain = append(ActivityDomain, value.toDomain())
	}

	return ActivityDomain, nil
}

func (cr *mysqlActivitiesRepository) GetByID(ctx context.Context, userId int) (activities.Domain, error) {
	rec := Activities{}
	err := cr.DB.Where("id = ?", userId).First(&rec).Error
	if err != nil {
		return activities.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (nr *mysqlActivitiesRepository) GetByName(ctx context.Context, ActivityName string) (activities.Domain, error) {
	rec := Activities{}
	err := nr.DB.Where("name = ?", ActivityName).First(&rec).Error
	if err != nil {
		return activities.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (nr *mysqlActivitiesRepository) Store(ctx context.Context, ActivitiesDomain *activities.Domain) (activities.Domain, error) {
	rec := fromDomain(ActivitiesDomain)

	result := nr.DB.Select("Name", "CreatedAt").Create(&rec)
	if result.Error != nil {
		return activities.Domain{}, result.Error
	}

	return rec.toDomain(), nil
}

func (nr *mysqlActivitiesRepository) Update(ctx context.Context, ActivitiesDomain *activities.Domain) (activities.Domain, error) {
	rec := fromDomain(ActivitiesDomain)

	result := nr.DB.Updates(rec)
	if result.Error != nil {
		return activities.Domain{}, result.Error
	}

	return rec.toDomain(), nil
}

func (nr *mysqlActivitiesRepository) Delete(ctx context.Context, ActivitiesDomain *activities.Domain) (activities.Domain, error) {
	rec := fromDomain(ActivitiesDomain)

	result := nr.DB.Delete(rec)

	if result.Error != nil {
		return activities.Domain{}, result.Error
	}

	return rec.toDomain(), nil
}

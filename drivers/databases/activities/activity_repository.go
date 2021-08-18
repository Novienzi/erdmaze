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

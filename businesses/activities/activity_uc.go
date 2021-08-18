package activities

import (
	"context"
	usecase "erdmaze/businesses"

	"time"
)

type activityUsecase struct {
	activityRepository Repository
	contextTimeout     time.Duration
}

func NewActivityUsecase(timeout time.Duration, cr Repository) Usecase {
	return &activityUsecase{
		contextTimeout:     timeout,
		activityRepository: cr,
	}
}

func (cu *activityUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	resp, err := cu.activityRepository.Find(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return resp, nil
}

func (cu *activityUsecase) GetByID(ctx context.Context, ActivityID int) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, cu.contextTimeout)
	defer cancel()

	if ActivityID <= 0 {
		return Domain{}, usecase.ErrActivityNotFound
	}
	res, err := cu.activityRepository.GetByID(ctx, ActivityID)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}

func (cu *activityUsecase) Store(ctx context.Context, ActivityDomain *Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, cu.contextTimeout)
	defer cancel()

	existedActivities, _ := cu.activityRepository.GetByName(ctx, ActivityDomain.Name)

	if existedActivities != (Domain{}) {
		return Domain{}, usecase.ErrDuplicateData
	}

	result, err := cu.activityRepository.Store(ctx, ActivityDomain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

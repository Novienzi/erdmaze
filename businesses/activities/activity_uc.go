package activities

import (
	"context"
	usecase "erdmaze/businesses"
	"strings"
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

func (cu *activityUsecase) Fetch(ctx context.Context, page, perpage int) ([]Domain, int, error) {
	ctx, cancel := context.WithTimeout(ctx, cu.contextTimeout)
	defer cancel()

	if page <= 0 {
		page = 1
	}
	if perpage <= 0 {
		perpage = 25
	}

	res, total, err := cu.activityRepository.Fetch(ctx, page, perpage)
	if err != nil {
		return []Domain{}, 0, err
	}

	return res, total, nil
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

func (cu *activityUsecase) GetByName(ctx context.Context, ActivityName string) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, cu.contextTimeout)
	defer cancel()

	if strings.TrimSpace(ActivityName) == "" {
		return Domain{}, usecase.ErrActivityNotFound
	}
	res, err := cu.activityRepository.GetByName(ctx, ActivityName)
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

func (cu *activityUsecase) Update(ctx context.Context, ActivitiesDomain *Domain) (*Domain, error) {
	existedActivities, err := cu.activityRepository.GetByID(ctx, ActivitiesDomain.ID)
	if err != nil {
		return &Domain{}, err
	}
	ActivitiesDomain.ID = existedActivities.ID

	result, err := cu.activityRepository.Update(ctx, ActivitiesDomain)
	if err != nil {
		return &Domain{}, err
	}

	return &result, nil
}

func (cu *activityUsecase) Delete(ctx context.Context, ActivitiesDomain *Domain) (*Domain, error) {
	existedActivities, err := cu.activityRepository.GetByID(ctx, ActivitiesDomain.ID)
	if err != nil {
		return &Domain{}, err
	}
	ActivitiesDomain.ID = existedActivities.ID

	result, err := cu.activityRepository.Delete(ctx, ActivitiesDomain)
	if err != nil {
		return &Domain{}, err
	}

	return &result, nil
}

package locations

import (
	"context"
	usecase "erdmaze/businesses"

	"time"
)

type locationUsecase struct {
	locationRepository Repository
	contextTimeout     time.Duration
}

func NewLocationUsecase(timeout time.Duration, cr Repository) Usecase {
	return &locationUsecase{
		contextTimeout:     timeout,
		locationRepository: cr,
	}
}

func (cu *locationUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	resp, err := cu.locationRepository.Find(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return resp, nil
}

func (cu *locationUsecase) GetByID(ctx context.Context, locationID int) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, cu.contextTimeout)
	defer cancel()

	if locationID <= 0 {
		return Domain{}, usecase.ErrLocationNotFound
	}
	res, err := cu.locationRepository.GetByID(ctx, locationID)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}

func (cu *locationUsecase) Store(ctx context.Context, locationDomain *Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, cu.contextTimeout)
	defer cancel()

	existedLocation, _ := cu.locationRepository.GetByName(ctx, locationDomain.Name)

	if existedLocation != (Domain{}) {
		return Domain{}, usecase.ErrDuplicateData
	}

	result, err := cu.locationRepository.Store(ctx, locationDomain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

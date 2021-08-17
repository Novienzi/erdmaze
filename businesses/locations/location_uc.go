package locations

import (
	"context"
	usecase "erdmaze/businesses"
	"strings"
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

func (cu *locationUsecase) Fetch(ctx context.Context, page, perpage int) ([]Domain, int, error) {
	ctx, cancel := context.WithTimeout(ctx, cu.contextTimeout)
	defer cancel()

	if page <= 0 {
		page = 1
	}
	if perpage <= 0 {
		perpage = 25
	}

	res, total, err := cu.locationRepository.Fetch(ctx, page, perpage)
	if err != nil {
		return []Domain{}, 0, err
	}

	return res, total, nil
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

func (cu *locationUsecase) GetByName(ctx context.Context, locationName string) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, cu.contextTimeout)
	defer cancel()

	if strings.TrimSpace(locationName) == "" {
		return Domain{}, usecase.ErrLocationNotFound
	}
	res, err := cu.locationRepository.GetByName(ctx, locationName)
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

func (cu *locationUsecase) Update(ctx context.Context, locationDomain *Domain) (*Domain, error) {
	existedActivities, err := cu.locationRepository.GetByID(ctx, locationDomain.ID)
	if err != nil {
		return &Domain{}, err
	}
	locationDomain.ID = existedActivities.ID

	result, err := cu.locationRepository.Update(ctx, locationDomain)
	if err != nil {
		return &Domain{}, err
	}

	return &result, nil
}

func (cu *locationUsecase) Delete(ctx context.Context, locationDomain *Domain) (*Domain, error) {
	existedActivities, err := cu.locationRepository.GetByID(ctx, locationDomain.ID)
	if err != nil {
		return &Domain{}, err
	}
	locationDomain.ID = existedActivities.ID

	result, err := cu.locationRepository.Delete(ctx, locationDomain)
	if err != nil {
		return &Domain{}, err
	}

	return &result, nil
}

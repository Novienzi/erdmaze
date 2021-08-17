package tourismpackages

import (
	"context"
	"erdmaze/businesses"
	"erdmaze/businesses/activities"
	"erdmaze/businesses/locations"
	"strings"
	"time"
)

type tourismPackagesUsecase struct {
	tourismPackagesRepository Repository
	activtyUsecase            activities.Usecase
	locationUsecase           locations.Usecase
	contextTimeout            time.Duration
}

func NewTourismPackagesUsecase(nr Repository, cu activities.Usecase, lu locations.Usecase, timeout time.Duration) Usecase {
	return &tourismPackagesUsecase{
		tourismPackagesRepository: nr,
		activtyUsecase:            cu,
		locationUsecase:           lu,
		contextTimeout:            timeout,
	}
}

func (nu *tourismPackagesUsecase) Fetch(ctx context.Context, page, perpage int) ([]Domain, int, error) {
	ctx, cancel := context.WithTimeout(ctx, nu.contextTimeout)
	defer cancel()

	if page <= 0 {
		page = 1
	}
	if perpage <= 0 {
		perpage = 25
	}

	res, total, err := nu.tourismPackagesRepository.Fetch(ctx, page, perpage)
	if err != nil {
		return []Domain{}, 0, err
	}

	return res, total, nil
}

func (nu *tourismPackagesUsecase) Store(ctx context.Context, tourismPackagesDomain *Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, nu.contextTimeout)
	defer cancel()

	_, err := nu.activtyUsecase.GetByID(ctx, tourismPackagesDomain.ActivityID)
	if err != nil {
		return Domain{}, businesses.ErrActivityNotFound
	}

	_, err = nu.locationUsecase.GetByID(ctx, tourismPackagesDomain.LocationID)
	if err != nil {
		return Domain{}, businesses.ErrLocationNotFound
	}

	existedNews, err := nu.tourismPackagesRepository.GetByName(ctx, tourismPackagesDomain.Name)
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return Domain{}, err
		}
	}
	if existedNews != (Domain{}) {
		return Domain{}, businesses.ErrDuplicateData
	}

	result, err := nu.tourismPackagesRepository.Store(ctx, tourismPackagesDomain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (nu *tourismPackagesUsecase) GetByID(ctx context.Context, tourismId int) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, nu.contextTimeout)
	defer cancel()

	if tourismId <= 0 {
		return Domain{}, businesses.ErrTourismsIDResource
	}
	res, err := nu.tourismPackagesRepository.GetByID(ctx, tourismId)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}

func (cu *tourismPackagesUsecase) GetAll(ctx context.Context, tourismName string, locationName string, activityName string) ([]Domain, error) {
	resp, err := cu.tourismPackagesRepository.GetAll(ctx, tourismName, locationName, activityName)
	if err != nil {
		return []Domain{}, err
	}
	return resp, nil
}

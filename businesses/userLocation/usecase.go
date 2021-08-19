package user_locations

import (
	"context"
	"encoding/json"
	"erdmaze/businesses"
	"erdmaze/businesses/iplocator"
	"log"
	"strings"
	"time"
)

type userLocationUsecase struct {
	userLocationRepository Repository
	contextTimeout         time.Duration
	ipLocator              iplocator.Repository
}

func NewUserLocationUsecase(nr Repository, timeout time.Duration, il iplocator.Repository) Usecase {
	return &userLocationUsecase{
		userLocationRepository: nr,
		contextTimeout:         timeout,
		ipLocator:              il,
	}
}

func (nu *userLocationUsecase) GetByUserID(ctx context.Context, userId int) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, nu.contextTimeout)
	defer cancel()

	if userId <= 0 {
		return Domain{}, businesses.ErrNewsIDResource
	}
	res, err := nu.userLocationRepository.GetByUserID(ctx, userId)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}

func (nu *userLocationUsecase) Store(ctx context.Context, ip string, userLocationDomain *Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, nu.contextTimeout)
	defer cancel()

	if strings.TrimSpace(ip) != "" {
		ipLoc, err := nu.ipLocator.GetLocationByIP(ctx, ip)
		if err != nil {
			log.Default().Printf("%+v", err)
		}
		jsonMarshal, err := json.Marshal(ipLoc)
		if err != nil {
			log.Default().Printf("%+v", err)
		}

		userLocationDomain.IPStat = string(jsonMarshal)
	}

	result, err := nu.userLocationRepository.Store(ctx, userLocationDomain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

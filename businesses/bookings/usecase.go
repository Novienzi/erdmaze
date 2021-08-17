package bookings

import (
	"context"
	usecase "erdmaze/businesses"
	"time"
)

type bookingsUsecase struct {
	bookingsRepository Repository
	contextTimeout     time.Duration
}

func NewBookingsUsecase(br Repository, timeout time.Duration) Usecase {
	return &bookingsUsecase{
		bookingsRepository: br,
		contextTimeout:     timeout,
	}
}

func (uc *bookingsUsecase) Fetch(ctx context.Context, page, perpage int) ([]Domain, int, int, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	if page <= 0 {
		page = 1
	}
	if perpage <= 0 {
		perpage = 10
	}

	res, total, err := uc.bookingsRepository.Fetch(ctx, page, perpage)
	lastPage := total / perpage

	if total%perpage > 0 {
		lastPage += 1
	}

	if err != nil {
		return []Domain{}, 0, 1, err
	}
	return res, total, lastPage, nil
}

func (uc *bookingsUsecase) GetByUserID(ctx context.Context, UserID int) ([]Domain, error) {
	resp, err := uc.bookingsRepository.GetByUserID(ctx, UserID)
	if err != nil {
		return []Domain{}, err
	}
	return resp, nil
}

func (uc *bookingsUsecase) GetByID(ctx context.Context, ID int) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	if ID <= 0 {
		return Domain{}, usecase.ErrBookingsIDResource
	}
	res, err := uc.bookingsRepository.GetByID(ctx, ID)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}

func (uc *bookingsUsecase) Store(ctx context.Context, bookingDomain *Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	existedFavourites, _ := uc.bookingsRepository.GetByID(ctx, bookingDomain.ID)

	if existedFavourites != (Domain{}) {
		return Domain{}, usecase.ErrDuplicateData
	}

	result, err := uc.bookingsRepository.Store(ctx, bookingDomain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

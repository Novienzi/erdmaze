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
		return Domain{}, usecase.ErrNotFound
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

	existedBooking, _ := uc.bookingsRepository.GetByID(ctx, bookingDomain.ID)

	if existedBooking != (Domain{}) {
		return Domain{}, usecase.ErrDuplicateData
	}

	result, err := uc.bookingsRepository.Store(ctx, bookingDomain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (uc *bookingsUsecase) Delete(ctx context.Context, bookingDomain *Domain) (*Domain, error) {
	existedBooking, err := uc.bookingsRepository.GetByID(ctx, bookingDomain.ID)
	if err != nil {
		return &Domain{}, err
	}
	bookingDomain.ID = existedBooking.ID

	result, err := uc.bookingsRepository.Delete(ctx, bookingDomain)
	if err != nil {
		return &Domain{}, err
	}

	return &result, nil
}

package users

import (
	"context"
	"erdmaze/app/middleware"
	"erdmaze/businesses"
	"erdmaze/helpers/encrypt"
	"strings"
	"time"
)

type userUsecase struct {
	userRepository Repository
	contextTimeout time.Duration
	jwtAuth        *middleware.ConfigJWT
}

func NewUserUsecase(ur Repository, jwtauth *middleware.ConfigJWT, timeout time.Duration) Usecase {
	return &userUsecase{
		userRepository: ur,
		jwtAuth:        jwtauth,
		contextTimeout: timeout,
	}
}

func (uc *userUsecase) Login(ctx context.Context, username, password string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	if strings.TrimSpace(username) == "" && strings.TrimSpace(password) == "" {
		return "", businesses.ErrUsernamePasswordNotFound
	}

	userDomain, err := uc.userRepository.GetByUsername(ctx, username)
	if err != nil {
		return "", err
	}

	if !encrypt.ValidateHash(password, userDomain.Password) {
		return "", businesses.ErrInternalServer
	}

	token := uc.jwtAuth.GenerateToken(userDomain.Id)
	return token, nil
}

func (uc *userUsecase) Store(ctx context.Context, userDomain *Domain) error {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	existedUser, err := uc.userRepository.GetByUsername(ctx, userDomain.Username)
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return err
		}
	}
	if existedUser != (Domain{}) {
		return businesses.ErrDuplicateData
	}

	userDomain.Password, err = encrypt.Hash(userDomain.Password)
	if err != nil {
		return businesses.ErrInternalServer
	}
	err = uc.userRepository.Store(ctx, userDomain)
	if err != nil {
		return err
	}

	return nil
}

func (uc *userUsecase) GetByID(ctx context.Context, userId int) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	if userId <= 0 {
		return Domain{}, businesses.ErrNewsIDResource
	}
	res, err := uc.userRepository.GetByID(ctx, userId)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}

func (cu *userUsecase) Update(ctx context.Context, usersDomain *Domain) (*Domain, error) {
	existedUsers, err := cu.userRepository.GetByID(ctx, usersDomain.Id)
	if err != nil {
		return &Domain{}, err
	}
	usersDomain.Id = existedUsers.Id

	result, err := cu.userRepository.Update(ctx, usersDomain)
	if err != nil {
		return &Domain{}, err
	}

	return &result, nil
}

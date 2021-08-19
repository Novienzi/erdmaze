package tourismpackages_test

import (
	"context"
	"erdmaze/businesses"
	activity "erdmaze/businesses/activities"
	location "erdmaze/businesses/locations"
	tourismPackageMock "erdmaze/businesses/tourism_packages/mocks"
	"errors"
	"os"
	"testing"

	tourismPackage "erdmaze/businesses/tourism_packages"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	tourismPackageRepository tourismPackageMock.Repository
	tourismPackageUsecase    tourismPackage.Usecase
	activtyUsecase           activity.Usecase
	locationUsecase          location.Usecase
)

func setup() {
	tourismPackageUsecase = tourismPackage.NewTourismPackagesUsecase(&tourismPackageRepository, activtyUsecase, locationUsecase, 2)
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())

}

func TestGetById(t *testing.T) {
	t.Run("test case 1, valid test ", func(t *testing.T) {
		domain := tourismPackage.Domain{
			ID:   1,
			Name: "hiking",
		}
		tourismPackageRepository.On("GetByID", context.Background(), mock.AnythingOfType("int")).Return(domain, nil).Once()

		result, err := tourismPackageRepository.GetByID(context.Background(), 1)

		assert.Nil(t, err)

		assert.Equal(t, domain.ID, result.ID)
		assert.Equal(t, domain.Name, result.Name)
	})

	t.Run("test case 2, invalid id", func(t *testing.T) {
		result, err := tourismPackageUsecase.GetByID(context.Background(), -1)

		assert.Equal(t, result, tourismPackage.Domain{})
		assert.Equal(t, err, businesses.ErrTourismsIDResource)
	})

}

func TestGetAll(t *testing.T) {
	t.Run("test case 1, get all", func(t *testing.T) {
		domain := []tourismPackage.Domain{
			{
				ID:   1,
				Name: "hiking",
			},
			{
				ID:   2,
				Name: "farming",
			},
		}
		tourismPackageRepository.On("Find", context.Background()).Return(domain, nil).Once()

		result, err := tourismPackageUsecase.GetAll(context.Background(), mock.Anything, mock.Anything, mock.Anything)

		assert.Equal(t, 2, len(result))
		assert.Nil(t, err)
	})

	t.Run("test case 2, repository error", func(t *testing.T) {
		errRepository := errors.New("mysql not running")
		tourismPackageRepository.On("Find", context.Background()).Return([]tourismPackage.Domain{}, errRepository).Once()

		result, err := tourismPackageUsecase.GetAll(context.Background(), mock.Anything, mock.Anything, mock.Anything)

		assert.Equal(t, 0, len(result))
		assert.Equal(t, errRepository, err)
	})
}

func TestStore(t *testing.T) {
	// t.Run("test case 1, valid test ", func(t *testing.T) {
	// 	domain := tourismPackage.Domain{
	// 		ID:   1,
	// 		Name: "hiking",
	// 	}
	// 	tourismPackageRepository.On("GetByName", mock.Anything, mock.AnythingOfType("string")).Return(tourismPackage.Domain{}, nil).Once()
	// 	tourismPackageRepository.On("Store", mock.Anything, mock.AnythingOfType("&tourismPackage.Domain")).Return(domain).Once()

	// 	result, err := tourismPackageUsecase.Store(context.Background(), &domain)

	// 	assert.NoError(t, err)
	// 	assert.Equal(t, domain.ID, result.ID)
	// 	assert.Equal(t, domain.Name, result.Name)
	// })

	t.Run("test case 2, duplicate data", func(t *testing.T) {
		domain := tourismPackage.Domain{
			ID:   1,
			Name: "hiking",
		}
		errRepository := errors.New("duplicate data")
		tourismPackageRepository.On("GetByName", mock.Anything, mock.AnythingOfType("string")).Return(domain, errRepository).Once()

		_, err := tourismPackageUsecase.Store(context.Background(), &domain)

		assert.Equal(t, err, businesses.ErrDuplicateData)
	})

	// t.Run("test case 3, store data failed", func(t *testing.T) {
	// 	domain := tourismPackage.Domain{
	// 		ID:   1,
	// 		Name: "hiking",
	// 	}
	// 	errRepository := errors.New("store data failed")
	// 	tourismPackageRepository.On("GetByName", mock.Anything, mock.AnythingOfType("string")).Return(tourismPackage.Domain{}, nil).Once()
	// 	tourismPackageRepository.On("Store", mock.Anything, mock.AnythingOfType("*tourismPackage.domain")).Return(errRepository).Once()

	// 	_, err := tourismPackageUsecase.Store(context.Background(), &domain)

	// 	assert.Equal(t, err, errRepository)
	// })

}

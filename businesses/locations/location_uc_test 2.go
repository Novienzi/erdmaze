package locations_test

import (
	"context"
	"erdmaze/businesses"
	locationMock "erdmaze/businesses/locations/mocks"
	"errors"
	"os"
	"testing"

	location "erdmaze/businesses/locations"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	locationRepository locationMock.Repository
	locationUsecase    location.Usecase
)

func setup() {
	locationUsecase = location.NewLocationUsecase(2, &locationRepository)
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())

}

func TestGetById(t *testing.T) {
	t.Run("test case 1, valid test ", func(t *testing.T) {
		domain := location.Domain{
			ID:   1,
			Name: "hiking",
		}
		locationRepository.On("GetByID", context.Background(), mock.AnythingOfType("int")).Return(domain, nil).Once()

		result, err := locationRepository.GetByID(context.Background(), 1)

		assert.Nil(t, err)

		assert.Equal(t, domain.ID, result.ID)
		assert.Equal(t, domain.Name, result.Name)
	})

	t.Run("test case 2, invalid id", func(t *testing.T) {
		result, err := locationUsecase.GetByID(context.Background(), -1)

		assert.Equal(t, result, location.Domain{})
		assert.Equal(t, err, businesses.ErrLocationNotFound)
	})

}

func TestGetAll(t *testing.T) {
	t.Run("test case 1, get all", func(t *testing.T) {
		domain := []location.Domain{
			{
				ID:   1,
				Name: "hiking",
			},
			{
				ID:   2,
				Name: "farming",
			},
		}
		locationRepository.On("Find", context.Background()).Return(domain, nil).Once()

		result, err := locationUsecase.GetAll(context.Background())

		assert.Equal(t, 2, len(result))
		assert.Nil(t, err)
	})

	t.Run("test case 2, repository error", func(t *testing.T) {
		errRepository := errors.New("mysql not running")
		locationRepository.On("Find", context.Background()).Return([]location.Domain{}, errRepository).Once()

		result, err := locationUsecase.GetAll(context.Background())

		assert.Equal(t, 0, len(result))
		assert.Equal(t, errRepository, err)
	})
}

func TestStore(t *testing.T) {
	// t.Run("test case 1, valid test ", func(t *testing.T) {
	// 	domain := location.Domain{
	// 		ID:   1,
	// 		Name: "hiking",
	// 	}
	// 	locationRepository.On("GetByName", mock.Anything, mock.AnythingOfType("string")).Return(location.Domain{}, nil).Once()
	// 	locationRepository.On("Store", mock.Anything, mock.AnythingOfType("&location.Domain")).Return(domain).Once()

	// 	result, err := locationUsecase.Store(context.Background(), &domain)

	// 	assert.NoError(t, err)
	// 	assert.Equal(t, domain.ID, result.ID)
	// 	assert.Equal(t, domain.Name, result.Name)
	// })

	t.Run("test case 2, duplicate data", func(t *testing.T) {
		domain := location.Domain{
			ID:   1,
			Name: "hiking",
		}
		errRepository := errors.New("duplicate data")
		locationRepository.On("GetByName", mock.Anything, mock.AnythingOfType("string")).Return(domain, errRepository).Once()

		_, err := locationUsecase.Store(context.Background(), &domain)

		assert.Equal(t, err, businesses.ErrDuplicateData)
	})

	// t.Run("test case 3, store data failed", func(t *testing.T) {
	// 	domain := location.Domain{
	// 		ID:   1,
	// 		Name: "hiking",
	// 	}
	// 	errRepository := errors.New("store data failed")
	// 	locationRepository.On("GetByName", mock.Anything, mock.AnythingOfType("string")).Return(location.Domain{}, nil).Once()
	// 	locationRepository.On("Store", mock.Anything, mock.AnythingOfType("*location.domain")).Return(errRepository).Once()

	// 	_, err := locationUsecase.Store(context.Background(), &domain)

	// 	assert.Equal(t, err, errRepository)
	// })

}

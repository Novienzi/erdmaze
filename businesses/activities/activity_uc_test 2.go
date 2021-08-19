package activities_test

import (
	"context"
	"erdmaze/businesses"
	activityMock "erdmaze/businesses/activities/mocks"
	"errors"
	"os"
	"testing"

	activity "erdmaze/businesses/activities"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	activityRepository activityMock.Repository
	activityUsecase    activity.Usecase
)

func setup() {
	activityUsecase = activity.NewActivityUsecase(2, &activityRepository)
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())

}

func TestGetById(t *testing.T) {
	t.Run("test case 1, valid test ", func(t *testing.T) {
		domain := activity.Domain{
			ID:   1,
			Name: "hiking",
		}
		activityRepository.On("GetByID", context.Background(), mock.AnythingOfType("int")).Return(domain, nil).Once()

		result, err := activityRepository.GetByID(context.Background(), 1)

		assert.Nil(t, err)

		assert.Equal(t, domain.ID, result.ID)
		assert.Equal(t, domain.Name, result.Name)
	})

	t.Run("test case 2, invalid id", func(t *testing.T) {
		result, err := activityUsecase.GetByID(context.Background(), -1)

		assert.Equal(t, result, activity.Domain{})
		assert.Equal(t, err, businesses.ErrActivityNotFound)
	})

}

func TestGetAll(t *testing.T) {
	t.Run("test case 1, get all", func(t *testing.T) {
		domain := []activity.Domain{
			{
				ID:   1,
				Name: "hiking",
			},
			{
				ID:   2,
				Name: "farming",
			},
		}
		activityRepository.On("Find", context.Background()).Return(domain, nil).Once()

		result, err := activityUsecase.GetAll(context.Background())

		assert.Equal(t, 2, len(result))
		assert.Nil(t, err)
	})

	t.Run("test case 2, repository error", func(t *testing.T) {
		errRepository := errors.New("mysql not running")
		activityRepository.On("Find", context.Background()).Return([]activity.Domain{}, errRepository).Once()

		result, err := activityUsecase.GetAll(context.Background())

		assert.Equal(t, 0, len(result))
		assert.Equal(t, errRepository, err)
	})
}

func TestStore(t *testing.T) {
	// t.Run("test case 1, valid test ", func(t *testing.T) {
	// 	domain := activity.Domain{
	// 		ID:   1,
	// 		Name: "hiking",
	// 	}
	// 	activityRepository.On("GetByName", mock.Anything, mock.AnythingOfType("string")).Return(activity.Domain{}, nil).Once()
	// 	activityRepository.On("Store", mock.Anything, mock.AnythingOfType("&activity.Domain")).Return(domain).Once()

	// 	result, err := activityUsecase.Store(context.Background(), &domain)

	// 	assert.NoError(t, err)
	// 	assert.Equal(t, domain.ID, result.ID)
	// 	assert.Equal(t, domain.Name, result.Name)
	// })

	t.Run("test case 2, duplicate data", func(t *testing.T) {
		domain := activity.Domain{
			ID:   1,
			Name: "hiking",
		}
		errRepository := errors.New("duplicate data")
		activityRepository.On("GetByName", mock.Anything, mock.AnythingOfType("string")).Return(domain, errRepository).Once()

		_, err := activityUsecase.Store(context.Background(), &domain)

		assert.Equal(t, err, businesses.ErrDuplicateData)
	})

	// t.Run("test case 3, store data failed", func(t *testing.T) {
	// 	domain := activity.Domain{
	// 		ID:   1,
	// 		Name: "hiking",
	// 	}
	// 	errRepository := errors.New("store data failed")
	// 	activityRepository.On("GetByName", mock.Anything, mock.AnythingOfType("string")).Return(activity.Domain{}, nil).Once()
	// 	activityRepository.On("Store", mock.Anything, mock.AnythingOfType("*activity.domain")).Return(errRepository).Once()

	// 	_, err := activityUsecase.Store(context.Background(), &domain)

	// 	assert.Equal(t, err, errRepository)
	// })

}

package usescase_test

import (
	"testing"

	"uala.com/core-service/internal/domain"
	"uala.com/core-service/internal/entity"
	"uala.com/core-service/internal/usescase"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockFollowRepository struct {
	mock.Mock
}

func (m *MockFollowRepository) Follow(follow entity.Follow) error {
	args := m.Called(follow)
	return args.Error(0)
}

func (m *MockFollowRepository) Unfollow(follow entity.Follow) error {
	args := m.Called(follow)
	return args.Error(0)
}

func TestFollowUserImpl_FollowUser(t *testing.T) {
	mockRepo := new(MockFollowRepository)
	followUserImpl := usescase.NewFollowUserImpl(mockRepo)

	follow := domain.Follow{
		FollowerID: 123,
		FolloweeID: 456,
	}

	followEntity := entity.Follow{
		FollowerID: 123,
		FolloweeID: 456,
	}

	mockRepo.On("Follow", followEntity).Return(nil)

	err := followUserImpl.FollowUser(follow)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestFollowUserImpl_UnfollowUser(t *testing.T) {
	mockRepo := new(MockFollowRepository)
	followUserImpl := usescase.NewFollowUserImpl(mockRepo)

	follow := domain.Follow{
		FollowerID: 123,
		FolloweeID: 456,
	}

	followEntity := entity.Follow{
		FollowerID: 123,
		FolloweeID: 456,
	}

	mockRepo.On("Unfollow", followEntity).Return(nil)

	err := followUserImpl.UnfollowUser(follow)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

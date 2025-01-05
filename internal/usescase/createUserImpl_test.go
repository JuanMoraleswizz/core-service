package usescase_test

import (
	"errors"
	"testing"

	"uala.com/core-service/internal/domain"
	"uala.com/core-service/internal/entity"
	"uala.com/core-service/internal/usescase"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUserRepository is a mock implementation of the UserRepository interface
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(user *entity.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func TestCreateUserImpl_CreateUserRepository(t *testing.T) {
	mockRepo := new(MockUserRepository)
	createUser := usescase.NewCreateUserImpl(mockRepo)

	t.Run("should return error if repository fails", func(t *testing.T) {
		user := domain.User{
			ID:       123,
			UserName: "testuser",
		}

		userEntity := &entity.User{
			ID:       user.ID,
			UserName: user.UserName,
		}

		mockRepo.On("Create", userEntity).Return(errors.New("repository error"))

		err := createUser.CreateUser(user)
		assert.EqualError(t, err, "repository error")
		mockRepo.AssertExpectations(t)
	})
}

func TestCreateUserImpl_CreateUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	createUser := usescase.NewCreateUserImpl(mockRepo)

	t.Run("success", func(t *testing.T) {
		user := domain.User{
			ID:       123,
			UserName: "testuser",
		}
		userEntity := &entity.User{
			ID:       user.ID,
			UserName: user.UserName,
		}

		mockRepo.On("Create", userEntity).Return(nil)

		err := createUser.CreateUser(user)
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})
}

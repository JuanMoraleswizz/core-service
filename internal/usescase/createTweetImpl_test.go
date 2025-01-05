package usescase

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"uala.com/core-service/internal/domain"
	"uala.com/core-service/internal/entity"
)

type MockTweetRepository struct {
	mock.Mock
}

func TestCreateTweetImpl_CreateTweetRepository(t *testing.T) {
	mockRepo := new(MockTweetRepository)
	createTweet := NewCreateTweetImpl(mockRepo)

	t.Run("should return error if repository fails", func(t *testing.T) {
		tweet := domain.Tweet{
			Content: "This is a valid tweet content.",
			UserID:  123,
		}

		tweetEntity := &entity.Tweet{
			Content: tweet.Content,
			UserID:  tweet.UserID,
		}

		mockRepo.On("Create", tweetEntity).Return(errors.New("repository error"))

		err := createTweet.CreateTweet(tweet)
		assert.EqualError(t, err, "repository error")
		mockRepo.AssertExpectations(t)
	})
}

func (m *MockTweetRepository) Create(tweet *entity.Tweet) error {
	args := m.Called(tweet)
	return args.Error(0)
}

func TestCreateTweetImpl_CreateTweet(t *testing.T) {
	mockRepo := new(MockTweetRepository)
	createTweet := NewCreateTweetImpl(mockRepo)

	t.Run("should return error if tweet content is too long", func(t *testing.T) {
		tweet := domain.Tweet{
			Content: "This is a very long tweet content that exceeds the 280 characters limit. " +
				"This is a very long tweet content that exceeds the 280 characters limit. " +
				"This is a very long tweet content that exceeds the 280 characters limit. " +
				"This is a very long tweet content that exceeds the 280 characters limit.",
			UserID: 123,
		}

		err := createTweet.CreateTweet(tweet)
		assert.EqualError(t, err, "tweet content is too long")
	})

	t.Run("should create tweet successfully", func(t *testing.T) {
		tweet := domain.Tweet{
			Content: "This is a valid tweet content.",
			UserID:  123,
		}

		tweetEntity := &entity.Tweet{
			Content: tweet.Content,
			UserID:  tweet.UserID,
		}

		mockRepo.On("Create", tweetEntity).Return(nil)

		err := createTweet.CreateTweet(tweet)
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

}

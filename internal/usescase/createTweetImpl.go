package usescase

import (
	"errors"

	"uala.com/core-service/internal/domain"
	"uala.com/core-service/internal/entity"
	"uala.com/core-service/internal/repository"
)

type CreateTweetImpl struct {
	TweetRepository repository.TweetRepository
}

func NewCreateTweetImpl(tweetRepository repository.TweetRepository) CreateTweet {
	return &CreateTweetImpl{TweetRepository: tweetRepository}
}

func (c *CreateTweetImpl) CreateTweet(tweet domain.Tweet) error {
	if len(tweet.Content) > 280 {
		return errors.New("tweet content is too long")
	}

	tweetEntity := &entity.Tweet{
		Content: tweet.Content,
		UserID:  tweet.UserID,
	}

	if err := c.TweetRepository.Create(tweetEntity); err != nil {
		return err
	}
	return nil
}

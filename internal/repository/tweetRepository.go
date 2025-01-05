package repository

import "uala.com/core-service/internal/entity"

type TweetRepository interface {
	Create(tweet *entity.Tweet) error
}

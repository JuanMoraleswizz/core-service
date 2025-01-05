package usescase

import "uala.com/core-service/internal/domain"

type CreateTweet interface {
	CreateTweet(tweet domain.Tweet) error
}

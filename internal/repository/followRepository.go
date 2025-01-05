package repository

import "uala.com/core-service/internal/entity"

type FollowRepository interface {
	Follow(follow entity.Follow) error
	Unfollow(follow entity.Follow) error
}

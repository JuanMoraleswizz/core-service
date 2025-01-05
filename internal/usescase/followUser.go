package usescase

import "uala.com/core-service/internal/domain"

type FollowUser interface {
	FollowUser(follow domain.Follow) error
	UnfollowUser(follow domain.Follow) error
}

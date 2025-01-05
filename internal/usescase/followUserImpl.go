package usescase

import (
	"uala.com/core-service/internal/domain"
	"uala.com/core-service/internal/entity"
	"uala.com/core-service/internal/repository"
)

type FollowUserImpl struct {
	FollowRepository repository.FollowRepository
}

func NewFollowUserImpl(followRepository repository.FollowRepository) FollowUser {
	return &FollowUserImpl{FollowRepository: followRepository}
}

func (f *FollowUserImpl) FollowUser(follow domain.Follow) error {

	followEntity := entity.Follow{
		FollowerID: follow.FollowerID,
		FolloweeID: follow.FolloweeID,
	}
	if err := f.FollowRepository.Follow(followEntity); err != nil {
		return err
	}
	return nil
}

func (f *FollowUserImpl) UnfollowUser(follow domain.Follow) error {

	followEntity := entity.Follow{
		FollowerID: follow.FollowerID,
		FolloweeID: follow.FolloweeID,
	}
	if err := f.FollowRepository.Unfollow(followEntity); err != nil {
		return err
	}
	return nil
}
